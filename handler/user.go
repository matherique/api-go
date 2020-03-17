package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matherique/api-golang/models"
)

type user struct{}

func withLogging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		log.Printf("%s - %s", r.URL, r.Method)
		f(w, r)
	}
}

func (u user) index(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	users, err := models.User.GetAll(models.User{})

	if err != nil {
		log.Printf("Internal error: %v", err)
		return nil, http.StatusInternalServerError, fmt.Errorf("unable to get all users")
	}

	return users, http.StatusOK, fmt.Errorf("unable to get all users")
}

func (u user) store(w http.ResponseWriter, r *http.Request) (interface{}, int,{
	return interface{}, 0, nil
}

//User user router
func User(srv *mux.Router) {
	u := user{}

	sb := srv.PathPrefix("/user").Subrouter()

	sb.HandleFunc("", responseHandler(u.index)).Methods("GET")
	sb.HandleFunc("/", responseHandler(u.index)).Methods("GET")
	sb.HandleFunc("/", responseHandler(u.store)).Methods("POST")

}
