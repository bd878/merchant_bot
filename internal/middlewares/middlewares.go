package middlewares

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"

	"github.com/bd878/merchant_bot/internal/logger"
)

type Middlewares struct {
	pool *pgxpool.Pool
	tableName string
}

func NewMiddlewares() *Middlewares {
	return &Middlewares{}
}

func HasMessageFromMiddleware(h bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, bot *bot.Bot, update *models.Update) {
		if update.Message != nil && update.Message.From != nil {
			h(ctx, bot, update)
		} else {
			logger.Log.Error("message.from is not given")
		}
	}
}

func AnswerCallbackQueryMiddleware(h bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		if update.CallbackQuery != nil {
			ok, err := b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
				CallbackQueryID: update.CallbackQuery.ID,
				ShowAlert:       false,
			})
			if !ok {
				logger.Log.Errorln("answer callback query returned false")
			}
			if err != nil {
				logger.Log.Errorw("failed ot answer callback query", "error", err)
			}
		}
		h(ctx, b, update)
	}
}