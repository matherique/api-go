package domain

import (
	"errors"
	"time"
)

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (u *User) Validade() error {
	if u.Name == "" {
		return errors.New("name required")
	}

	if u.Email == "" {
		return errors.New("email required")
	}

	return nil
}
