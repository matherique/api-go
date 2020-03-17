package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

// Response of an user
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

// ResponseHandler - handle response with log
func responseHandler(h func(w http.ResponseWriter, r *http.Request) (interface{}, int, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// log 
		log.Printf("%s - %s", r.URL, r.Method)

		data, status, err := h(w, r)

		if err != nil {
			data = err.Error()
		}

		w.WriteHeader(status)
		w.Header().Set("Content-Type", "application/j son")

		err = json.NewEncoder(w).Encode(Response{Data: data, Success: err != nil})

		if err != nil {
			log.Printf("could not encode response to output: %v", err)
		}
	}
}
