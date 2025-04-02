package clients

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/go-telegram/bot/models"
)

func (m Module) RestoreChat(ctx context.Context, chat *models.Chat) error {
	_, ok := m.app.Chats().Get(chat.ID)
	if !ok {
		_, err := m.repo.FindChat(ctx, chat.ID)
		if err != nil {
			if err == pgx.ErrNoRows {
				err = m.repo.SaveChat(ctx, chat)
				if err != nil {
					return err
				}
			} else {
				return err
			}
		}

		m.app.Chats().Set(chat.ID, chat)
	}
	return nil
}