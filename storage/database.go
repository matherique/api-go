package storage

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

type Database struct {
	ctx *context.Context
}

func (d *Database) Connect() (*sql.DB, error) {
	return sql.Open("postgres", "host=database user=postgres dbname=api password=123 sslmode=verify-full")
}
