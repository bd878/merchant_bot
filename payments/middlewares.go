package payments

import (
	"strings"
	"strconv"
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type (
	idKey struct {}
)

func (m Module) GetTransactionIDMiddleware(h bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		if update.CallbackQuery != nil {
			m.log.Debugw("transaction id", "data", update.CallbackQuery.Data)
			parts := strings.Split(update.CallbackQuery.Data, ":")
			transactionID := parts[len(parts)-1]
			id, err := strconv.Atoi(transactionID)
			if err != nil {
				m.log.Errorw("cannot parse transaction id", "id", id, "error", err)
				return
			}
			m.log.Debugw("id", "id", id)
			ctx = context.WithValue(ctx, &idKey{}, id)
		}
		h(ctx, b, update)
	}
}
