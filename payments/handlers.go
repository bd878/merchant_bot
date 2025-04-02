package payments

import (
	"fmt"
	"context"
	"github.com/google/uuid"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	merchant "github.com/bd878/merchant_bot"
)

var (
	xtr string = "XTR"
	invoice_payload string = "invoice payload"
)

func (m Module) InvoiceHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendInvoice(ctx, &bot.SendInvoiceParams{
		ChatID: update.Message.Chat.ID,
		Title: "Test gift",
		Description: "Test gift description",
		Payload: invoice_payload,
		Prices: []models.LabeledPrice{models.LabeledPrice{Label: xtr, Amount: 10}},
		Currency: xtr,
	})
	if err != nil {
		m.log.Errorw("send invoice returns error", "error", err)
	}
}

func (m Module) RefundTransactionhandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	id, ok := ctx.Value(&idKey{}).(int)
	if !ok {
		m.log.Errorw("no id key", "id", id)
		return
	}

	err := m.repo.RefundPayment(ctx, uint32(id))
	if err != nil {
		m.log.Errorw("failed to make a refund", "id", id, "error", err)
		return
	}

	payment, err := m.repo.FindPayment(ctx, uint32(id))
	if err != nil {
		m.log.Errorw("failed to find payment after refund", "id", id, "error", err)
		return
	}

	ok, err = b.RefundStarPayment(ctx, &bot.RefundStarPaymentParams{
		UserID: payment.UserID,
		TelegramPaymentChargeID: payment.TelegramPaymentChargeID,
	})
	if err != nil {
		m.log.Errorw("tg failed to refund star payment", "user_id", payment.UserID, "payment_id", payment.ID, "error", err)
		return
	}

	if !ok {
		m.log.Errorw("tg refund was not successful", "user_id", payment.UserID, "payment_id", payment.ID)
		return
	}

	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: payment.UserID,
		Text: fmt.Sprintf("%s", merchant.LangRu.Text("refunded_success")),
	})
	if err != nil {
		m.log.Errorw("failed to send refund message", "id", id, "error", err)
		return
	}
}

func (m Module) ShowTransactionHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	id, ok := ctx.Value(&idKey{}).(int)
	m.log.Debugw("transaction handler id", "id", id)
	if !ok {
		m.log.Errorw("no id key", "id", id)
		return
	}

	payment, err := m.repo.FindPayment(ctx, uint32(id))
	if err != nil {
		m.log.Errorw("cannot find transaction", "id", id, "error", err)
		return
	}

	if payment.Refunded {
		_, err = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: payment.UserID,
			Text: fmt.Sprintf("%d\n%s", payment.TotalAmount, merchant.LangRu.Text("refunded")),
		})
	} else {
		_, err = b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: payment.UserID,
			Text: fmt.Sprintf("%d", payment.TotalAmount),
			ReplyMarkup: RefundKeyboard(merchant.LangRu, payment.ID),
		})
	}

	if err != nil {
		m.log.Errorw("failed to send message", "id", id, "error", err)
		return
	}
}

func (m Module) PreCheckoutUpdateHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	ok, err := b.AnswerPreCheckoutQuery(ctx, &bot.AnswerPreCheckoutQueryParams{
		PreCheckoutQueryID: update.PreCheckoutQuery.ID,
		OK: true,
	})
	if err != nil {
		m.log.Errorw("failed to answer pre checkout query", "error", err)
	}
	if ok {
		m.log.Infoln("pre checkout query ok")
	} else {
		m.log.Warnln("pre checkout query is NOT ok")
	}
}

func (m Module) SuccessfullPaymentHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	id := uuid.New().ID()

	err := m.repo.SavePayment(ctx, &merchant.Payment{
		SuccessfulPayment: update.Message.SuccessfulPayment,
		ID: id,
		UserID: update.Message.From.ID,
	})

	if err != nil {
		m.log.Errorw("failed to save successfull payment", "error", err)
	}
}

func (m Module) ShowTransactions(ctx context.Context, b *bot.Bot, update *models.Update) {
	transactions, err := m.repo.ListUserTransactions(ctx, update.Message.From.ID, 10, 0)
	if err != nil {
		m.log.Errorw("failed to get user star transactions", "user_id", update.Message.From.ID, "error", err)
		return
	}

	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text: merchant.LangRu.Text("transactions"),
		ReplyMarkup: TransactionsKeyboard(merchant.LangRu, transactions),
	})
	if err != nil {
		m.log.Errorw("failed to send transactions", "error", err)
	}
}
