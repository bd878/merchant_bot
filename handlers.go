package merchant_bot

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	merchant "github.com/bd878/merchant_bot/models"
)

var (
	xtr string = "XTR"
	invoice_payload string = "invoice payload"
)

func (m ClientsModule) StartHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text: "ok",
	})
	if err != nil {
		log.Errorw("send message returns error", "error", err)
	}
}

func (m ClientsModule) InvoiceHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendInvoice(ctx, &bot.SendInvoiceParams{
		ChatID: update.Message.Chat.ID,
		Title: "Test gift",
		Description: "Test gift description",
		Payload: invoice_payload,
		Prices: []models.LabeledPrice{models.LabeledPrice{Label: xtr, Amount: 10}},
		Currency: xtr,
	})
	if err != nil {
		log.Errorw("send invoice returns error", "error", err)
	}
}

func (m ClientsModule) PreCheckoutUpdateHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	ok, err := b.AnswerPreCheckoutQuery(ctx, &bot.AnswerPreCheckoutQueryParams{
		PreCheckoutQueryID: update.PreCheckoutQuery.ID,
		OK: true,
	})
	if err != nil {
		log.Errorw("failed to answer pre checkout query", "error", err)
	}
	if ok {
		log.Infoln("pre checkout query ok")
	} else {
		log.Warnln("pre checkout query is NOT ok")
	}
}

func (m ClientsModule) SuccessfullPaymentHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	err := m.app.Repo().SavePayment(ctx, &merchant.Payment{
		SuccessfulPayment: update.Message.SuccessfulPayment,
		UserID: update.Message.From.ID,
	})
	if err != nil {
		log.Errorw("failed to save successfull payment", "error", err)
	}
}

func (m ClientsModule) MemberKickedHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	log.Infow("kicked", "chat_id", update.Message.Chat.ID)
}

func (m ClientsModule) MemberRestoredHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	log.Infow("restored", "chat_id", update.Message.Chat.ID)
}

// IMPORTANT for turn backs: https://core.telegram.org/bots/payments-stars#live-checklist
func (m ClientsModule) TermsHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text: LangRu.Text("terms"),
	})
	if err != nil {
		log.Errorw("failed to send terms", "error", err)
	}
}

func (m ClientsModule) ShowTransactions(ctx context.Context, b *bot.Bot, update *models.Update) {
	transactions, err := b.GetStarTransactions(ctx, &bot.GetStarTransactionsParams{
		Offset: 0,
		Limit: 50,
	})
	if err != nil {
		log.Errorw("failed to get star transactions", "error", err)
		return
	}

	_, err = b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text: LangRu.Text("transactions"),
		ReplyMarkup: TransactionsKeyboard(LangRu, transactions),
	})
	if err != nil {
		log.Errorw("failed to send transactions", "error", err)
	}
}