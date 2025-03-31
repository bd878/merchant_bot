package merchant_bot

import (
	"context"
)

type Monolith interface {
	Repo() *Repo
	Bot() *Bot
	Log() *Logger
	Config() Config
	Chats() *Chats
}

type Module interface {
	Startup(ctx context.Context, mono Monolith) error
}