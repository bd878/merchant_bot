package merchant_bot

import (
	"context"
	"github.com/go-telegram/bot"
)

type ClientsModule struct {
	app Monolith
}

func (m *ClientsModule) Startup(ctx context.Context, app Monolith) error {
	m.app = app

	app.Bot().RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, m.StartHandler, m.RestoreChatMiddleware)
	app.Bot().RegisterHandler(bot.HandlerTypeMessageText, "/invoice", bot.MatchTypeExact, m.InvoiceHandler, m.RestoreChatMiddleware)
	app.Bot().RegisterHandler(bot.HandlerTypeMessageText, "/transactions", bot.MatchTypeExact, m.ShowTransactions, m.RestoreChatMiddleware)
	app.Bot().RegisterHandler(bot.HandlerTypeMessageText, "/terms", bot.MatchTypeExact, m.TermsHandler, m.RestoreChatMiddleware)

	// TODO: receive inline invoices

	app.Bot().RegisterHandlerMatchFunc(PreCheckoutUpdateMatch, m.PreCheckoutUpdateHandler)
	app.Bot().RegisterHandlerMatchFunc(SuccessfullPaymentMatch, m.SuccessfullPaymentHandler, m.HasUserMiddleware)
	app.Bot().RegisterHandlerMatchFunc(MemberKickedMatch, m.MemberKickedHandler)
	app.Bot().RegisterHandlerMatchFunc(MemberRestoredMatch, m.MemberRestoredHandler)

	return nil
}