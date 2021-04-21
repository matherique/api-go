package routes

import (
	"net/http"
)

type homeRoutes struct{}

func (h homeRoutes) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GET /home"))
}

func LoadRoutes(server *http.ServeMux) {
	server.Handle("/", homeRoutes{})
	server.Handle("/users", usersRoute{})
}
