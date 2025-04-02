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