package models

import (
	"database/sql"
	// mysql drive
	_ "github.com/go-sql-driver/mysql"
)

// Connection database connection
func Connection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/smartpark")
	if err != nil {
		return nil, err
	}

	return db, nil
}
