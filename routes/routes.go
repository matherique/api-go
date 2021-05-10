package routes

import (
	"net/http"

	"github.com/matherique/api-go/storage"
	"github.com/matherique/api-go/utils"
)

func LoadRoutes(server *http.ServeMux) {
	database := storage.Database{}
	logger := utils.NewLogger()

	connection, err := database.Connect()

	if err != nil {
		logger.Fatalf(err.Error())
	}

	ur := NewUserRoute(connection, logger)
	hr := NewHomeRoute(logger)

	server.Handle("/", hr)
	server.Handle("/users", ur)
	server.Handle("/users/", ur)
}
