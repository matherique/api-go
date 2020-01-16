package models

// User that we create
type User struct {
	ID       string `json:"id,omitempty"`
	Nome     string `json:"nome,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

// Insert  user in the database
func (u User) Insert() {
	con := Connection()

	defer con.Close()

	insert, err := con.Query("INSERT INTO smartpark (nome, email, password) VALUES(?, ?, ?)", u.Nome, u.Email, u.Password)

	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
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
