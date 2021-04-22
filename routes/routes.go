package routes

import (
	"database/sql"
	"net/http"

	"github.com/matherique/api-go/utils"
)

func LoadRoutes(server *http.ServeMux) {
	var database *sql.DB

	ur := NewUserRoute(database, utils.NewLogger())
	hr := NewHomeRoute(utils.NewLogger())

	server.Handle("/", hr)
	server.Handle("/users", ur)
	server.Handle("/users/", ur)
}
