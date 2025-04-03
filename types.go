package merchant_bot

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ChatKey struct {}

type Monolith interface {
	Pool() *pgxpool.Pool
	Bot() *Bot
	Log() *Logger
	Config() Config
	Chats() *Chats
	History() *History
	Modules() []Module
}

type Module interface {
	Startup(ctx context.Context, mono Monolith) error
	Name() string
}