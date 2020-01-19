package models

import "golang.org/x/crypto/bcrypt"

// User that we create
type User struct {
	ID       string `json:"id,omitempty"`
	Nome     string `json:"nome,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

// Insert user in the database
func (u User) Insert() User {
	con := Connection()

	defer con.Close()

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	if err != nil {
		panic(err.Error())
	}

	result, err := con.Exec("INSERT INTO user (nome, email, password) VALUES(?, ?, ?)", u.Nome, u.Email, hash)

	if err != nil {
		panic(err.Error())

	}
	_, err = result.RowsAffected()

	if err != nil {
		panic(err.Error())
	}

	u.Password = string(hash)

	return u
}

// GetAll list all users in the database
func (u User) GetAll() []User {
	con := Connection()

	defer con.Close()
	results, err := con.Query("SELECT id, nome, email, password FROM user")

	if err != nil {
		panic(err.Error())
	}

	var users []User

	for results.Next() {
		var user User
		err = results.Scan(&user.ID, &user.Nome, &user.Email, &user.Password)

		if err != nil {
			panic(err.Error())
		}
		users = append(users, user)
	}

	return users

}

// CheckLogin verify if user and password exist
func (u User) CheckLogin() (bool, error) {

	con := Connection()

	defer con.Close()
	results, err := con.Query("SELECT password FROM user WHERE email = ?", u.Email)

	if err != nil {
		return false, err
	}

	var resUser User
	results.Next()
	err = results.Scan(&resUser.Password)

	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(resUser.Password), []byte(u.Password))

	if err != nil {
		return false, err
	}

	return true, nil
}
