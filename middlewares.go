package merchant_bot

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func HasUserMiddleware(h bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, bot *bot.Bot, update *models.Update) {
		if update.Message != nil && update.Message.From != nil {
			h(ctx, bot, update)
		} else {
			log.Error("message.from is not given")
		}
	}
}