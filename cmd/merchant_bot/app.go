package main

import (
	"sync"
	"context"
	"net/http"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/bd878/merchant_bot/internal/config"
	"github.com/bd878/merchant_bot/internal/chats"
	"github.com/bd878/merchant_bot/internal/bot"
	"github.com/bd878/merchant_bot/internal/logger"
	"github.com/bd878/merchant_bot/internal/system"
	"github.com/bd878/merchant_bot/internal/history"
)

type app struct {
	conf   config.Config
	pool  *pgxpool.Pool
	chats *chats.Chats
	bot   *bot.Bot
	log   *logger.Logger
	history *history.History
	modules []system.Module
}

func (a app) Config() config.Config {
	return a.conf
}

func (a app) Pool() *pgxpool.Pool {
	return a.pool
}

func (a app) Bot() *bot.Bot {
	return a.bot
}

func (a app) Chats() *chats.Chats {
	return a.chats
}

func (a app) Log() *logger.Logger {
	return a.log
}

func (a app) History() *history.History {
	return a.history
}

func (a app) Modules() []system.Module {
	return a.modules
}

func (a *app) startupModules() error {
	for _, m := range a.modules {
		if err := m.Startup(context.Background(), a); err != nil {
			return err
		}
	}
	return nil
}

func (a *app) waitForWeb(wg *sync.WaitGroup) {
	defer wg.Done()

	server := &http.Server{
		Addr: a.Config().Addr,
	}

	a.Log().Infow("starting http webhook server", "addr", a.Config().Addr)

	mux := http.NewServeMux()

	mux.HandleFunc(a.Config().WebhookPath, a.Bot().WebhookHandler())

	server.Handler = mux

	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		a.Log().Errorw("http server returned error", "error", err)
	}
}

func (a *app) waitForWebhook(wg *sync.WaitGroup) {
	defer wg.Done()

	a.Bot().StartWebhook(context.Background())
}