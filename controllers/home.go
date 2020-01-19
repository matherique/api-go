package controllers

import (
	"encoding/json"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/matherique/api-golang/models"
)

// HomeIndex page handler
func HomeIndex(w http.ResponseWriter, r *http.Request) {
	home := struct {
		Page string `json:"page"`
	}{Page: "Home page"}

	json.NewEncoder(w).Encode(home)
}

type loginReponse struct {
	Error string `json:"error,omitempty"`
	Token string `json:"token,omitempty"`
}

// Login page
func Login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var user models.User
	err := decoder.Decode(&user)

	if err != nil {
		panic(err)
	}

	valid, err := user.CheckLogin()
	if err != nil {
		resp := loginReponse{
			Error: err.Error(), Token: "",
		}
		json.NewEncoder(w).Encode(resp)
	}

	if valid == false {
		resp := loginReponse{
			Error: "invalid user/password", Token: "",
		}
		json.NewEncoder(w).Encode(resp)
	}

	// Create the Claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		ExpiresAt: 15000,
		Issuer:    "test",
	})

	tokenString, err := token.SignedString([]byte("secret"))

	if err != nil {
		resp := loginReponse{
			Error: err.Error(), Token: "",
		}
		json.NewEncoder(w).Encode(resp)
	}

	resp := loginReponse{
		Token: tokenString,
	}

	json.NewEncoder(w).Encode(resp)
}
