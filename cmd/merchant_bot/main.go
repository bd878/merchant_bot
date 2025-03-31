package main

import (
	"database/sql"
	"os"
	"sync"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/go-telegram/bot"

	merchant "github.com/bd878/merchant_bot"
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

	db, err := sql.Open("pgx", m.conf.PGConn)
	if err != nil {
		panic(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)

	m.repo = merchant.NewRepo("marchandise", db)

	m.modules = []merchant.Module{
		&merchant.ClientsModule{},
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