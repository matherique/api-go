package main

import (
	"encoding/json"
	"net/http"
)

type home struct{}

func (h home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	resposta := struct {
		Mensagem string
	}{
		Mensagem: "Ola mundo",
	}

	json.NewEncoder(w).Encode(resposta)
}
