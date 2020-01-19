package models

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

var (
	ctx context.Context
)

// User that we create
type User struct {
	ID       string `json:"id,omitempty"`
	Nome     string `json:"nome,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

// Insert user in the database
func (u User) Insert() (User, error) {
	con, err := Connection()

	if err != nil {
		return User{}, err
	}

	defer con.Close()

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		return User{}, err
	}

	stmt, err := con.PrepareContext(ctx, "INSERT INTO user (nome, email, password) VALUES(?, ?, ?)")

	if err != nil {
		return User{}, err
	}

	u.Password = string(hash)

	defer stmt.Close()

	if _, err := stmt.Exec(u.Nome, u.Email, u.Password); err != nil {
		return User{}, err
	}

	return u, nil
}

// GetAll list all users in the database
func (u User) GetAll() ([]User, error) {
	con, err := Connection()

	if err != nil {
		return []User{}, err
	}

	defer con.Close()

	rows, err := con.QueryContext(ctx, "SELECT id, nome, email, password FROM user")
	defer rows.Close()

	if err != nil {
		return []User{}, err
	}

	users := make([]User, 0)

	for rows.Next() {
		var user User

		if err = rows.Scan(&user.ID, &user.Nome, &user.Email, &user.Password); err != nil {
			return []User{}, err
		}

		users = append(users, user)
	}

	return users, nil

}

// CheckLogin verify if user and password exist
func (u User) CheckLogin() (bool, error) {

	con, err := Connection()

	if err != nil {
		return false, err
	}

	defer con.Close()

	var hashedPass string

	err = con.QueryRowContext(ctx, "SELECT password FROM user WHERE email = ?", u.Email).Scan(&hashedPass)

	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(u.Password))

	if err != nil {
		return false, err
	}

	return true, nil
}

// Update user
func (u User) Update() (User, error) {
	con, err := Connection()

	if err != nil {
		return User{}, err
	}

	defer con.Close()

	if u.Password != "" {
		hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

		if err != nil {
			panic(err.Error())
		}

		u.Password = string(hashed)
	}

	result, err := con.Exec("UPDATE INTO user $FIELD VALUES(?, ?, ?)", u.Nome, u.Email, u.Password)

	if err != nil {
		panic(err.Error())

	}
	_, err = result.RowsAffected()

	return u, err
}
