package clients

import (
	"context"
	"strings"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	merchant "github.com/bd878/merchant_bot"
)

type (
	langKey struct {}
)

func (m Module) RestoreChatMiddleware(h bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, bot *bot.Bot, update *models.Update) {
		if update.Message != nil {
			chat := &merchant.Chat{Chat: &update.Message.Chat, Lang: merchant.LangRu}
			err := m.RestoreChat(ctx, chat)
			if err != nil {
				m.log.Errorw("failed to restore a chat", "error", err)
				return
			}
			ctx = context.WithValue(ctx, &merchant.ChatKey{}, chat)
		}
		h(ctx, bot, update)
	}
}

func (m Module) LangMiddleware(h bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, bot *bot.Bot, update *models.Update) {
		if update.Message != nil {
			parts := strings.Split(update.CallbackQuery.Data, ":")
			langStr := parts[len(parts)-1]
			if langStr == "" {
				m.log.Errorw("empty lang string", "id", update.Message.Chat.ID)
				return
			}
			ctx = context.WithValue(ctx, &langKey{}, merchant.LangFromString(langStr))
		}
		h(ctx, bot, update)
	}
}