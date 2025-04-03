package clients

import (
	"fmt"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/bd878/merchant_bot/internal/pkg"
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

func (r Repository) Update(ctx context.Context, chat *pkg.Chat) error {
	const query = "UPDATE %s SET lang = $1 WHERE id = $2"

	_, err := r.pool.Exec(ctx, r.table(query), chat.Lang.String(), chat.ID)

	return err
}

func (r Repository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}