package domain

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) Normalize() {

}

func (u User) Validade() (bool, error) {
	return true, nil
}
