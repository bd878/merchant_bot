package merchant_bot

import (
	"fmt"
	"context"
	"database/sql"

	"github.com/go-telegram/bot/models"
)

type Repository struct {
	db *sql.DB
	tableName string
}

func NewRepository(tableName string, db *sql.DB) *Repository {
	return &Repository{
		db: db,
		tableName: tableName,
	}
}

func (r Repository) Find(ctx context.Context, chatID int64) (*models.Chat, error) {
const query = "SELECT type, title, username, first_name, last_name, is_forum FROM %s WHERE id = :id LIMIT 1"

	chat := &models.Chat{
		ID: chatID,
	}

	var chatType string

	err := r.db.QueryRowContext(ctx, r.table(query), sql.Named("id", chatID)).Scan(&chatType, &chat.Title,
		&chat.Username, &chat.FirstName, &chat.LastName, &chat.IsForum)
	if err != nil {
		return nil, err
	}

	chat.Type = models.ChatType(chatType)

	return chat, nil
}

func (r Repository) Save(ctx context.Context, chat *models.Chat) error {
	const query = "INSERT INTO %s (id, type, title, username, first_name, last_name, is_forum) VALUES (:id, :type, :title, :username, :firstName, :lastName, :isForum)"

	_, err := r.db.ExecContext(ctx, r.table(query), sql.Named("id", chat.ID), sql.Named("type", chat.Type), sql.Named("title", chat.Title),
		sql.Named("username", chat.Username), sql.Named("firstName", chat.FirstName), sql.Named("lastName", chat.LastName), sql.Named("isForum", chat.IsForum))

	return err
}

func (r Repository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}