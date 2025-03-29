package merchant_bot

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

var (
	xtr string = "XTR"
	invoice_payload string = "invoice payload"
)

func StartHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text: "ok",
	})
}

func InvoiceHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendInvoice(ctx, &bot.SendInvoiceParams{
		ChatID: update.Message.Chat.ID,
		Title: "Test gift",
		Description: "Test gift description",
		Payload: invoice_payload,
		Prices: []models.LabeledPrice{models.LabeledPrice{Label: xtr, Amount: 10}},
		Currency: xtr,
	})
}

func PreCheckoutUpdateHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.AnswerPreCheckoutQuery(ctx, &bot.AnswerPreCheckoutQueryParams{
		PreCheckoutQueryID: update.PreCheckoutQuery.ID,
		OK: true,
	})
}

func SuccessfullPaymentHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	log.Info(update.Message.SuccessfulPayment.TelegramPaymentChargeID)
}

func MemberKickedHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	log.Infow("kicked", "chat_id", update.Message.Chat.ID)
}

func MemberRestoredHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	log.Infow("restored", "chat_id", update.Message.Chat.ID)
}

// IMPORTANT for turn backs: https://core.telegram.org/bots/payments-stars#live-checklist
func TermsHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
}

func ShowTransactions(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.GetStarTransactions(ctx, &bot.GetStarTransactionsParams{
		Offset: 0,
		Limit: 50,
	})
}