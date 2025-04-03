package clients

import (
	"fmt"
	"context"

	"github.com/go-telegram/bot/models"
	"github.com/jackc/pgx/v5/pgxpool"

	merchant "github.com/bd878/merchant_bot"
)

type Repository struct {
	pool *pgxpool.Pool
	tableName string
}

func NewRepository(tableName string, pool *pgxpool.Pool) *Repository {
	return &Repository{
		pool: pool,
		tableName: tableName,
	}
}

func (r Repository) FindChat(ctx context.Context, chatID int64) (*merchant.Chat, error) {
	const query = "SELECT lang, type, title, username, first_name, last_name, is_forum FROM %s WHERE id = $1 LIMIT 1"

	chat := &merchant.Chat{
		Chat: &models.Chat{
			ID: chatID,
		},
	}

	var (
		chatType, lang string
	)

	err := r.pool.QueryRow(ctx, r.table(query), chatID).Scan(&lang, &chatType, &chat.Title,
		&chat.Username, &chat.FirstName, &chat.LastName, &chat.IsForum)
	if err != nil {
		return nil, err
	}

	chat.Type = models.ChatType(chatType)
	chat.Lang = merchant.LangFromString(lang)

	return chat, nil
}

func (r Repository) SaveChat(ctx context.Context, chat *merchant.Chat) error {
	const query = "INSERT INTO %s (lang, id, type, title, username, first_name, last_name, is_forum) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"

	_, err := r.pool.Exec(ctx, r.table(query), chat.Lang.String(), chat.ID, chat.Type, chat.Title,
		chat.Username, chat.FirstName, chat.LastName, chat.IsForum)

	return err
}

func (r Repository) Update(ctx context.Context, chat *merchant.Chat) error {
	const query = "UPDATE %s SET lang = $1 WHERE id = $2"

	_, err := r.pool.Exec(ctx, r.table(query), chat.Lang.String(), chat.ID)

	return err
}

func (r Repository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}