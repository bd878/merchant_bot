package system

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/bd878/merchant_bot/internal/config"
	"github.com/bd878/merchant_bot/internal/logger"
	"github.com/bd878/merchant_bot/internal/chats"
	"github.com/bd878/merchant_bot/internal/history"
	"github.com/bd878/merchant_bot/internal/bot"
)

type Monolith interface {
	Pool() *pgxpool.Pool
	Bot() *bot.Bot
	Log() *logger.Logger
	Config() config.Config
	Chats() *chats.Chats
	History() *history.History
	Modules() []Module
}

type Module interface {
	Startup(ctx context.Context, mono Monolith) error
	Name() string
}