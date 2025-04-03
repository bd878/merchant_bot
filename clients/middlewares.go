package clients

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	merchant "github.com/bd878/merchant_bot"
)

type chatKey struct {}

func (m Module) RestoreChatMiddleware(h bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, bot *bot.Bot, update *models.Update) {
		if update.Message != nil {
			chat := &merchant.Chat{Chat: &update.Message.Chat, Lang: merchant.LangRu}
			err := m.RestoreChat(ctx, chat)
			if err != nil {
				m.log.Errorw("failed to restore a chat", "error", err)
				return
			}
			ctx = context.WithValue(ctx, &chatKey{}, chat)
		}
		h(ctx, bot, update)
	}
}