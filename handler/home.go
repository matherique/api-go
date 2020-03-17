package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// // HomeIndex page handler
// func HomeIndex(w http.ResponseWriter, r *http.Request) {
// 	home := struct {
// 		Page string `json:"page"`
// 	}{Page: "Home page"}

// 	json.NewEncoder(w).Encode(home)
// }

// type loginReponse struct {
// 	Error string `json:"error,omitempty"`
// 	Token string `json:"token,omitempty"`
// }

// // Login page
// func Login(w http.ResponseWriter, r *http.Request) {
// 	decoder := json.NewDecoder(r.Body)

// 	var user models.User
// 	err := decoder.Decode(&user)

// 	if err != nil {
// 		panic(err)
// 	}

// 	valid, err := user.CheckLogin()
// 	if err != nil {
// 		resp := loginReponse{
// 			Error: err.Error(), Token: "",
// 		}
// 		json.NewEncoder(w).Encode(resp)
// 	}

// 	if valid == false {
// 		resp := loginReponse{
// 			Error: "invalid user/password", Token: "",
// 		}
// 		json.NewEncoder(w).Encode(resp)
// 	}

// 	// Create the Claims
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
// 		ExpiresAt: 15000,
// 		Issuer:    "test",
// 	})

// 	tokenString, err := token.SignedString([]byte("secret"))

// 	if err != nil {
// 		resp := loginReponse{
// 			Error: err.Error(), Token: "",
// 		}
// 		json.NewEncoder(w).Encode(resp)
// 	}

// 	resp := loginReponse{
// 		Token: tokenString,
// 	}

// }

type home struct{}

func (h home) index(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	fmt.Println("home")
	return "asdasd", 0, nil
}

//Home user router
func Home(srv *mux.Router) {
	h := home{}

	srv.HandleFunc("", responseHandler(h.index)).Methods("GET")
	srv.HandleFunc("/", responseHandler(h.index)).Methods("GET")
}
