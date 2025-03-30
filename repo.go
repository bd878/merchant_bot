package merchant_bot

import (
	"database/sql"
)

type Repo struct {
	pool *sql.DB
	dbName string
}

func NewRepo(dbName string, pool *sql.DB) *Repo {
	return &Repo{
		pool: pool,
		dbName: dbName,
	}
}