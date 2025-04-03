package main

import (
	"context"
	"os"
	"sync"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/go-telegram/bot"

	"github.com/bd878/merchant_bot/clients"
	"github.com/bd878/merchant_bot/payments"
	"github.com/bd878/merchant_bot/internal/config"
	"github.com/bd878/merchant_bot/internal/system"
	"github.com/bd878/merchant_bot/internal/chats"
	"github.com/bd878/merchant_bot/internal/history"
	"github.com/bd878/merchant_bot/internal/logger"
	merchantBot "github.com/bd878/merchant_bot/internal/bot"
)

func main() {
	configPath, ok := os.LookupEnv("CONFIG_FILE")
	if !ok {
		fmt.Printf("Usage: env CONFIG_FILE=<file> %s\n", os.Args[0])
		os.Exit(1)
	}

	m := app{}

	m.conf = config.LoadConfig(configPath)
	m.log = logger.NewLog()
	defer m.log.Sync()

	var err error
	m.pool, err = pgxpool.New(context.Background(), m.conf.PGConn)
	if err != nil {
		panic(err)
	}
	defer m.pool.Close()

	m.chats = chats.NewChats("marchandise.chat.chat", m.pool)
	m.history = history.NewHistory()

	m.bot = merchantBot.NewBot(os.Getenv("TELEGRAM_MERCHANT_BOT_TOKEN"),
		os.Getenv("TELEGRAM_MERCHANT_BOT_WEBHOOK_SECRET_TOKEN"), m.Config().WebhookURL + m.Config().WebhookPath,
		bot.WithDebug(),
		bot.WithMiddlewares(m.chats.RestoreChatMiddleware))

	// TODO: use grpc for inter-module communications
	m.modules = []system.Module{
		&clients.Module{},
		&payments.Module{},
	}

	if err := m.startupModules(); err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go m.waitForWebhook(&wg)
	wg.Add(1)
	go m.waitForWeb(&wg)
	wg.Wait()
}