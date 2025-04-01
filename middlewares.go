package merchant_bot

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type chatKey struct {}

func (m ClientsModule) RestoreChatMiddleware(h bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, bot *bot.Bot, update *models.Update) {
		if update.Message != nil {
			chat, ok := m.app.Chats().Get(update.Message.Chat.ID)
			if !ok {
				var err error
				chat, err = m.app.Repo().Find(ctx, update.Message.Chat.ID)
				if err != nil {
					if err == pgx.ErrNoRows {
						err = m.app.Repo().Save(ctx, &update.Message.Chat)
						if err != nil {
							log.Errorw("failed to save chat", "chat_id", update.Message.Chat.ID, "error", err)
							return
						}
					} else {
						log.Errorw("failed to find chat", "chat_id", update.Message.Chat.ID, "error", err)
						return
					}
				}

				m.app.Chats().Set(update.Message.Chat.ID, &update.Message.Chat)
			}

			ctx = context.WithValue(ctx, &chatKey{}, chat)
		}
		h(ctx, bot, update)
	}
}
