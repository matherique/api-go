package routes

import (
	"encoding/json"
	"errors"
	"net/http"
)

const (
	MethodGET  = "GET"
	MethodPOST = "POST"
	MethodPUT  = "PUT"
)

var (
	Error400 = errors.New("bad request")
	Error401 = errors.New("unauthorized")
	Error404 = errors.New("not found")
	Error500 = errors.New("internal server error")
)

func handleError(f func(http.ResponseWriter, *http.Request) error) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)

		if err == nil {
			return
		}

		switch err {
		case Error400:
			badRequest(w, r)
		case Error401:
			unauthorized(w, r)
		case Error404:
			notFound(w, r)
		case Error500:
			internalServerError(w, r)
		default:
			internalServerError(w, r)
		}
	}
}

func created(w http.ResponseWriter, r *http.Request, response interface{}) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func ok(w http.ResponseWriter, r *http.Request, response interface{}) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("not found"))
}

func internalServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("internal server error"))
}

func unauthorized(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("unauthorized"))
}

func badRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("bad request"))
}
