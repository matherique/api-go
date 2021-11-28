package storage

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

type Db interface {
	Connect() (*sql.DB, error)
}

func NewDb() Db {
	db := new(database)

	return db
}

type database struct {
	ctx *context.Context
}

func (d *database) Connect() (*sql.DB, error) {
	db, _ := sql.Open("postgres", "postgres://postgres:123@localhost/apidb?sslmode=disable")

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
