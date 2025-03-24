package merchant_bot

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func StartHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text: "ok",
	})
}

func MemberKickedHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	log.Infow("kicked", "chat_id", update.Message.Chat.ID)
}

func MemberRestoredHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	log.Infow("restored", "chat_id", update.Message.Chat.ID)
}