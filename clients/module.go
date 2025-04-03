package clients

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/bd878/merchant_bot/internal/logger"
	"github.com/bd878/merchant_bot/internal/system"
)

type Module struct {
	repo *Repository
	log *logger.Logger
	app system.Monolith
}

func (m *Module) Startup(ctx context.Context, app system.Monolith) error {
	m.repo = NewRepository("marchandise.chat.chat", app.Pool())
	m.log = app.Log()
	m.app = app

	app.Bot().RegisterHandler(bot.HandlerTypeCallbackQueryData, "ru:", bot.MatchTypePrefix, m.ChangeLanguageHandler, m.LangMiddleware)
	app.Bot().RegisterHandler(bot.HandlerTypeCallbackQueryData, "en:", bot.MatchTypePrefix, m.ChangeLanguageHandler, m.LangMiddleware)
	app.Bot().RegisterHandlerMatchFunc(MemberKickedMatch, m.MemberKickedHandler)
	app.Bot().RegisterHandlerMatchFunc(MemberRestoredMatch, m.MemberRestoredHandler)

	return nil
}

func (Module) Name() string { return "clients" }

func MemberKickedMatch(update *models.Update) bool {
	if update.MyChatMember != nil {
		return (
			update.MyChatMember.NewChatMember.Type == models.ChatMemberTypeBanned ||
			update.MyChatMember.NewChatMember.Type == models.ChatMemberTypeLeft)
	}
	return false
}

func MemberRestoredMatch(update *models.Update) bool {
	if update.MyChatMember != nil {
		return (
			update.MyChatMember.NewChatMember.Type == models.ChatMemberTypeMember ||
			update.MyChatMember.NewChatMember.Type == models.ChatMemberTypeAdministrator)
	}
	return false
}
