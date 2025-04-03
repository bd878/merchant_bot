package clients

import (
	"context"
	"github.com/jackc/pgx/v5"

	merchant "github.com/bd878/merchant_bot"
)

func (m Module) RestoreChat(ctx context.Context, chat *merchant.Chat) error {
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

func (m Module) Find(ctx context.Context, id int64) (*merchant.Chat, error) {
	chat, ok := m.app.Chats().Get(id)
	if !ok {
		chat, err := m.repo.FindChat(ctx, id)
		if err != nil {
			return nil, err
		}
		return chat, nil
	}
	return chat, nil
}