package merchant_bot

import (
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type Bot struct {
	*bot.Bot
}

func must(token string, opts ...bot.Option) *bot.Bot {
	b, err := bot.New(token, opts...)
	if err != nil {
		panic(err)
	}
	return b
}

func NewBot(token, webhookToken, webhookURL string, opts ...bot.Option) *Bot {
	opts = append(opts, bot.WithDefaultHandler(defaultHandler))
	b := &Bot{must(token, opts...)}

	b.SetWebhook(context.Background(), &bot.SetWebhookParams{
		URL:      webhookURL,
		SecretToken: webhookToken,
	})

	return b
}

func defaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		log.Errorln("message is nil, exit")
		return
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "pong",
	})
}