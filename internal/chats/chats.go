package chats

import (
	"sync"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"

	"github.com/bd878/merchant_bot/internal/pkg"
	"github.com/bd878/merchant_bot/internal/i18n"
)

type Chats struct {
	mu sync.Mutex
	dict map[int64]*pkg.Chat
	repo *Repository
}

func NewChats(tableName string, pool *pgxpool.Pool) *Chats {
	repo := NewRepository(tableName, pool)
	return &Chats{
		dict: make(map[int64]*pkg.Chat, 0),
		repo: repo,
	}
}

func (c *Chats) RestoreChatMiddleware(h bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, bot *bot.Bot, update *models.Update) {
		if update.Message != nil {
			chat := &pkg.Chat{Chat: &update.Message.Chat, Lang: i18n.LangRu}
			_, ok := c.Get(chat.ID)
			if !ok {
				_, err := c.repo.FindChat(ctx, chat.ID)
				if err != nil {
					if err == pgx.ErrNoRows {
						err = c.repo.CreateChat(ctx, chat)
						if err != nil {
							return
						}
					} else {
						return
					}
				}

				c.Set(chat.ID, chat)
			}
			ctx = context.WithValue(ctx, &pkg.ChatKey{}, chat)
		}
		h(ctx, bot, update)
	}
}

func (c Chats) Get(id int64) (*pkg.Chat, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	chat, ok := c.dict[id]
	if !ok {
		return nil, false
	}
	return chat, true
} 

func (c *Chats) Set(id int64, chat *pkg.Chat) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.dict[id] = chat
}