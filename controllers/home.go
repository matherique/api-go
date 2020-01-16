package controllers

import (
	"encoding/json"
	"net/http"
)

// HomeIndex page handler
func HomeIndex(w http.ResponseWriter, r *http.Request) {
	home := struct {
		page string
	}{page: "Home page"}

	json.NewEncoder(w).Encode(home)
}
