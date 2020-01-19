package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/matherique/api-golang/models"
)

// Response of an user
type Response struct {
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}

// UserIndex page handler
func UserIndex(w http.ResponseWriter, r *http.Request) {
	alluser, err := models.User.GetAll(models.User{})
	if err != nil {
		json.NewEncoder(w).Encode(Response{Error: err.Error()})
	}
	json.NewEncoder(w).Encode(Response{Data: alluser})
}

// UserStore page handler
func UserStore(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var user models.User
	err := decoder.Decode(&user)

	if err != nil {
		json.NewEncoder(w).Encode(Response{Error: err.Error()})
		return
	}

	insertedUser, err := user.Insert()

	if err != nil {
		json.NewEncoder(w).Encode(Response{Error: err.Error()})
		return
	}

	json.NewEncoder(w).Encode(Response{Data: insertedUser})
}

// UserUpdate page handler
func UserUpdate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user models.User
	err := decoder.Decode(&user)

	if err != nil {
		json.NewEncoder(w).Encode(Response{Error: err.Error()})
		return
	}

}
