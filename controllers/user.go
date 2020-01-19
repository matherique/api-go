package controllers

import (
	"encoding/json"
	"github.com/matherique/api-golang/models"
	"net/http"
)

// Response of an user
type Response struct {
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}

// UserIndex page handler
func UserIndex(w http.ResponseWriter, r *http.Request) {
	alluser := models.User.GetAll(models.User{})

	json.NewEncoder(w).Encode(Response{Data: alluser})
}

// UserStore page handler
func UserStore(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var user models.User
	err := decoder.Decode(&user)

	if err != nil {
		panic(err)
	}

	insertedUser := user.Insert()
	json.NewEncoder(w).Encode(Response{Data: insertedUser})
}
