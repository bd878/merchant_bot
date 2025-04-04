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
	"github.com/bd878/merchant_bot/internal/logger"
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
		var (
			chat *models.Chat
			value *pkg.Chat
			ok bool
			err error
		)

		if update.Message != nil {
			chat = &update.Message.Chat
		} else if update.CallbackQuery != nil {
			chat = &update.CallbackQuery.Message.Message.Chat
		}

		if chat != nil {
			value, ok = c.Get(chat.ID)
			if !ok {
				value, err = c.repo.FindChat(ctx, chat.ID)
				if err != nil {
					if err == pgx.ErrNoRows {
						value = &pkg.Chat{Chat: chat, Lang: i18n.LangRu}
						err = c.repo.CreateChat(ctx, value)
						if err != nil {
							logger.Log.Errorw("failed to create chat", "error", err)
							return
						}
					} else {
						logger.Log.Errorw("failed to find chat", "error", err)
						return
					}
				}

				c.Set(chat.ID, value)
			}
			ctx = context.WithValue(ctx, &pkg.ChatKey{}, value)
			h(ctx, bot, update)
		} else {
			logger.Log.Errorln("cannot restore chat middlware, no chat")
		}
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