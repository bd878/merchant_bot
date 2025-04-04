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
		if update.CallbackQuery != nil {
			parts := strings.Split(update.CallbackQuery.Data, ":")
			langStr := parts[0]
			if langStr == "" {
				m.log.Errorw("empty lang string", "id", update.CallbackQuery.From.ID)
				return
			}
			m.log.Debugw("select lang", "lang", langStr)
			ctx = context.WithValue(ctx, &pkg.LangKey{}, i18n.LangFromString(langStr))
		} else {
			m.log.Warnln("update is not a CallbackQuery")
		}
		h(ctx, bot, update)
	}
}