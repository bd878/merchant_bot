package merchant_bot

import (
	"fmt"
	"context"

	"github.com/go-telegram/bot/models"
	"github.com/jackc/pgx/v5/pgxpool"
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

func (r Repository) Find(ctx context.Context, chatID int64) (*models.Chat, error) {
	const query = "SELECT type, title, username, first_name, last_name, is_forum FROM %s WHERE id = $1 LIMIT 1"

	chat := &models.Chat{
		ID: chatID,
	}

	var chatType string

	err := r.pool.QueryRow(ctx, r.table(query), chatID).Scan(&chatType, &chat.Title,
		&chat.Username, &chat.FirstName, &chat.LastName, &chat.IsForum)
	if err != nil {
		return nil, err
	}

	chat.Type = models.ChatType(chatType)

	return chat, nil
}

func (r Repository) Save(ctx context.Context, chat *models.Chat) error {
	const query = "INSERT INTO %s (id, type, title, username, first_name, last_name, is_forum) VALUES ($1, $2, $3, $4, $5, $6, $7)"

	_, err := r.pool.Exec(ctx, r.table(query), chat.ID, chat.Type, chat.Title, chat.Username, chat.FirstName, chat.LastName, chat.IsForum)

	return err
}

func (r Repository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}