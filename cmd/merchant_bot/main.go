package main

import (
	"context"
	"os"
	"sync"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/go-telegram/bot"

	merchant "github.com/bd878/merchant_bot"
	"github.com/bd878/merchant_bot/clients"
	"github.com/bd878/merchant_bot/payments"
)

func main() {
	configPath, ok := os.LookupEnv("CONFIG_FILE")
	if !ok {
		fmt.Printf("Usage: env CONFIG_FILE=<file> %s\n", os.Args[0])
		os.Exit(1)
	}

	m := app{}

	m.conf = merchant.LoadConfig(configPath)
	m.log = merchant.NewLog()
	defer m.log.Sync()

	m.bot = merchant.NewBot(os.Getenv("TELEGRAM_MERCHANT_BOT_TOKEN"),
		os.Getenv("TELEGRAM_MERCHANT_BOT_WEBHOOK_SECRET_TOKEN"), m.Config().WebhookURL + m.Config().WebhookPath,
		bot.WithDebug())

	var err error
	m.pool, err = pgxpool.New(context.Background(), m.conf.PGConn)
	if err != nil {
		panic(err)
	}
	defer m.pool.Close()

	m.chats = merchant.NewChats()

	m.modules = []merchant.Module{
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