package merchant_bot

import (
	"sync"
	"github.com/go-telegram/bot/models"
)

type Chats struct {
	mu sync.Mutex
	dict map[int64]*models.Chat
}

func NewChats() *Chats {
	return &Chats{
		dict: make(map[int64]*models.Chat, 0),
	}
}

func (c Chats) Get(id int64) (*models.Chat, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	chat, ok := c.dict[id]
	if !ok {
		return nil, false
	}
	return chat, true
} 

func (c *Chats) Set(id int64, chat *models.Chat) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.dict[id] = chat
}