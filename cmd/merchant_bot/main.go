package main

import (
	"os/signal"
	"os"
	"sync"
	"fmt"
	"context"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	merchant "github.com/bd878/merchant_bot"
)

func main() {
	configPath, ok := os.LookupEnv("CONFIG_FILE")
	if !ok {
		fmt.Printf("Usage: env CONFIG_FILE=<file> %s\n", os.Args[0])
		os.Exit(1)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	defer merchant.Log().Sync()

	conf := merchant.LoadConfig(configPath)

	opts := []bot.Option{
		bot.WithDefaultHandler(defaultHandler),
	}

	b, err := bot.New(os.Getenv("TELEGRAM_MERCHANT_BOT_TOKEN"), opts...)
	if err != nil {
		merchant.Log().Fatalw("cannot create bot", "error", err)
	}

	merchant.SetBot(b)

	b.SetWebhook(ctx, &bot.SetWebhookParams{
		URL:      conf.WebhookURL + conf.WebhookPath,
		SecretToken: os.Getenv("TELEGRAM_MERCHANT_BOT_WEBHOOK_SECRET_TOKEN"),
	})

	httpServer := merchant.NewHTTPServer(conf.Addr, conf.WebhookPath, b.WebhookHandler())

	var wg sync.WaitGroup
	wg.Add(1)
	go httpServer.Run(&wg)
	wg.Wait()
}

func defaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		merchant.Log().Errorln("message is nil, exit")
		return
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "pong",
	})
}