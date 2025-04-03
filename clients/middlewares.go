package clients

import (
	"context"
	"strings"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/bd878/merchant_bot/internal/i18n"
	"github.com/bd878/merchant_bot/internal/pkg"
)

func (m Module) LangMiddleware(h bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, bot *bot.Bot, update *models.Update) {
		if update.Message != nil {
			parts := strings.Split(update.CallbackQuery.Data, ":")
			langStr := parts[len(parts)-1]
			if langStr == "" {
				m.log.Errorw("empty lang string", "id", update.Message.Chat.ID)
				return
			}
			ctx = context.WithValue(ctx, &pkg.LangKey{}, i18n.LangFromString(langStr))
		}
		h(ctx, bot, update)
	}
}