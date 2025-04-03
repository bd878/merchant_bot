package payments

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/bd878/merchant_bot/internal/logger"
	"github.com/bd878/merchant_bot/internal/system"
	"github.com/bd878/merchant_bot/internal/middlewares"
)

type Module struct {
	repo *Repository
	log *logger.Logger
}

func (m *Module) Startup(ctx context.Context, app system.Monolith) error {
	m.repo = NewRepository("marchandise.payments.payments", app.Pool())
	m.log = app.Log()

	app.Bot().RegisterHandler(bot.HandlerTypeMessageText, "/invoice", bot.MatchTypeExact, m.InvoiceHandler)
	app.Bot().RegisterHandler(bot.HandlerTypeMessageText, "/transactions", bot.MatchTypeExact, m.ShowTransactions, middlewares.HasMessageFromMiddleware)

	app.Bot().RegisterHandler(bot.HandlerTypeCallbackQueryData, "tr:", bot.MatchTypePrefix, m.ShowTransactionHandler, middlewares.AnswerCallbackQueryMiddleware, m.GetTransactionIDMiddleware)
	app.Bot().RegisterHandler(bot.HandlerTypeCallbackQueryData, "refund:", bot.MatchTypePrefix, m.RefundTransactionhandler, middlewares.AnswerCallbackQueryMiddleware, m.GetTransactionIDMiddleware)
	app.Bot().RegisterHandlerMatchFunc(PreCheckoutUpdateMatch, m.PreCheckoutUpdateHandler)
	app.Bot().RegisterHandlerMatchFunc(SuccessfullPaymentMatch, m.SuccessfullPaymentHandler, middlewares.HasMessageFromMiddleware)

	return nil
}

func (Module) Name() string { return "payments" }

func PreCheckoutUpdateMatch(update *models.Update) bool {
	return update.PreCheckoutQuery != nil
}

func SuccessfullPaymentMatch(update *models.Update) bool {
	if update.Message != nil {
		return update.Message.SuccessfulPayment != nil
	}
	return false
}