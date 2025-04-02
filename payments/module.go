package payments

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	merchant "github.com/bd878/merchant_bot"
)

type Module struct {
	repo *Repository
	log *merchant.Logger
}

func (m *Module) Startup(ctx context.Context, app merchant.Monolith) error {
	m.repo = NewRepository("marchandise.payments.payments", app.Pool())
	m.log = app.Log()

	app.Bot().RegisterHandler(bot.HandlerTypeMessageText, "/invoice", bot.MatchTypeExact, m.InvoiceHandler)
	app.Bot().RegisterHandler(bot.HandlerTypeMessageText, "/transactions", bot.MatchTypeExact, m.ShowTransactions, merchant.HasUserMiddleware)

	app.Bot().RegisterHandler(bot.HandlerTypeCallbackQueryData, "tr:", bot.MatchTypePrefix, m.ShowTransactionHandler)
	app.Bot().RegisterHandlerMatchFunc(PreCheckoutUpdateMatch, m.PreCheckoutUpdateHandler)
	app.Bot().RegisterHandlerMatchFunc(SuccessfullPaymentMatch, m.SuccessfullPaymentHandler, merchant.HasUserMiddleware)

	return nil
}

func PreCheckoutUpdateMatch(update *models.Update) bool {
	return update.PreCheckoutQuery != nil
}

func SuccessfullPaymentMatch(update *models.Update) bool {
	if update.Message != nil {
		return update.Message.SuccessfulPayment != nil
	}
	return false
}