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
	db, err := sql.Open("postgres", "postgres://postgres:123@localhost/apidb?sslmode=disable")

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, err
}
