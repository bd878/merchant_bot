package main

import (
	"sync"
	"context"
	"net/http"
	merchant "github.com/bd878/merchant_bot"
)

type app struct {
	conf   merchant.Config
	repo  *merchant.Repo
	chats *merchant.Chats
	bot   *merchant.Bot
	log   *merchant.Logger
	modules []merchant.Module
}

func (a *app) Config() merchant.Config {
	return a.conf
}

func (a *app) Repo() *merchant.Repo {
	return a.repo
}

func (a *app) Bot() *merchant.Bot {
	return a.bot
}

func (a *app) Chats() *merchant.Chats {
	return a.chats
}

func (a *app) Log() *merchant.Logger {
	return a.log
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