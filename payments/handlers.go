package payments

import (
	"fmt"
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	merchant "github.com/bd878/merchant_bot"
	merchantModels "github.com/bd878/merchant_bot/models"
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
	err := m.repo.SavePayment(ctx, &merchantModels.Payment{
		SuccessfulPayment: update.Message.SuccessfulPayment,
		UserID: update.Message.From.ID,
	})
	if err != nil {
		m.log.Errorw("failed to save successfull payment", "error", err)
	}
}

func (m Module) ShowTransactions(ctx context.Context, b *bot.Bot, update *models.Update) {
	transactions, err := b.GetStarTransactions(ctx, &bot.GetStarTransactionsParams{
		Offset: 0,
		Limit: 50,
	})
	if err != nil {
		m.log.Errorw("failed to get star transactions", "error", err)
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


func TransactionsKeyboard(code merchant.LangCode, trans *models.StarTransactions) *models.InlineKeyboardMarkup {
	keyboards := make([][]models.InlineKeyboardButton, 0)
	for _, tr := range trans.Transactions {
		keyboards = append(keyboards, []models.InlineKeyboardButton{
			{Text: fmt.Sprintf("%d - %d", tr.Amount, tr.Date), CallbackData: "trans"},
		})
	}
	return &models.InlineKeyboardMarkup{
		InlineKeyboard: keyboards,
	}
}