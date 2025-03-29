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
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text: "ok",
	})
	if err != nil {
		log.Errorw("send message returns error", "error", err)
	}
}

func InvoiceHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
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

func PreCheckoutUpdateHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	ok, err := b.AnswerPreCheckoutQuery(ctx, &bot.AnswerPreCheckoutQueryParams{
		PreCheckoutQueryID: update.PreCheckoutQuery.ID,
		OK: true,
	})
	if err != nil {
		log.Errorw("failed to ansewer pre checkout query", "error", err)
	}
	if ok {
		log.Infoln("pre checkout query ok")
	} else {
		log.Warnln("pre checkout query is NOT ok")
	}
}

func SuccessfullPaymentHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	log.Info("payment success", update.Message.SuccessfulPayment.TelegramPaymentChargeID)
}

func MemberKickedHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	log.Infow("kicked", "chat_id", update.Message.Chat.ID)
}

func MemberRestoredHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	log.Infow("restored", "chat_id", update.Message.Chat.ID)
}

// IMPORTANT for turn backs: https://core.telegram.org/bots/payments-stars#live-checklist
func TermsHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text: LangRu.Text("terms"),
	})
	if err != nil {
		log.Errorw("failed to send terms", "error", err)
	}
}

func ShowTransactions(ctx context.Context, b *bot.Bot, update *models.Update) {
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