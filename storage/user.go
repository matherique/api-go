package models

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var (
	ctx context.Context
)

// User that we create
type User struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
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

	stmt, err := con.Prepare("INSERT INTO user (name, email, password) VALUES(?, ?, ?)")

	if err != nil {
		return User{}, err
	}

	u.Password = string(hash)

	defer stmt.Close()

	if _, err := stmt.Exec(u.Name, u.Email, u.Password); err != nil {
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

	rows, err := con.Query("SELECT id, name, email FROM user")
	defer rows.Close()

	if err != nil {
		return []User{}, err
	}

	users := make([]User, 0)

	for rows.Next() {
		var user User

		if err = rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
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

	err = con.QueryRow("SELECT password FROM user WHERE email = ?", u.Email).Scan(&hashedPass)

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
func (u User) Update(id string) (bool, error) {
	con, err := Connection()

	if err != nil {
		return false, err
	}

	defer con.Close()

	if u.Password != "" {
		hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

		if err != nil {
			return false, err
		}

		u.Password = string(hashed)
	}

	parsed := strucToMap(&u)

	sets := make([]string, len(parsed))
	values := make([]interface{}, len(parsed))

	for k, v := range parsed {
		if v != "" {
			sets = append(sets, fmt.Sprintf("%s = ?", k))
			values = append(values, v)
		}
	}

	values = append(values, id)

	setformated := strings.Join(sets, " and ")
	fmt.Println(setformated)
	fmt.Println(sets)
	querystring := fmt.Sprintf("UPDATE balances SET %s WHERE user_id = ?", setformated)

	result, err := con.Exec(querystring, values...)

	if err != nil {
		panic(err.Error())

	}
	rows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	if rows != 1 {
		return false, nil
	}

	return true, nil
}

func strucToMap(item interface{}) map[string]interface{} {
	var result map[string]interface{}

	jsonfydata, err := json.Marshal(item)

	if err != nil {
		log.Printf("[Error]: %s", err.Error())
	}

	err = json.Unmarshal(jsonfydata, &result)

	if err != nil {
		log.Printf("[Error]: %s", err.Error())
	}

	return result
}
