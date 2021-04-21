package routes

import (
	"encoding/json"
	"net/http"
)

const (
	MethodGET  = "GET"
	MethodPOST = "POST"
	MethodPUT  = "PUT"
)

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

func badRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("bad request"))
}
