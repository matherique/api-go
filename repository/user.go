package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/matherique/api-go/domain"
)

type UserRepository struct {
	Database *sql.DB
}

func (repo *UserRepository) FindAll() ([]domain.User, error) {
	users := make([]domain.User, 0)

	rows, err := repo.Database.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var email string
		var username string
		var password string
		var createdAt time.Time
		var updatedAt time.Time

		if err := rows.Scan(
			&id,
			&name,
			&email,
			&username,
			&password,
			&createdAt,
			&updatedAt,
		); err != nil {
			return nil, err
		}

		user := domain.User{
			Id:        id,
			Name:      name,
			Email:     email,
			Username:  username,
			Password:  password,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}

		users = append(users, user)
	}

	return users, nil
}

func (repo *UserRepository) Create(user domain.User, ctx context.Context) ([]domain.User, error) {
	return []domain.User{}, nil
}
