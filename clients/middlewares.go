package clients

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func (m Module) RestoreChatMiddleware(h bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, bot *bot.Bot, update *models.Update) {
		if update.Message != nil {
			err := m.RestoreChat(ctx, &update.Message.Chat)
			if err != nil {
				m.log.Errorw("failed to restore a chat", "error", err)
				return
			}
		}
		h(ctx, bot, update)
	}
}