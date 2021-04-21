package routes

import (
	"net/http"
	"regexp"
)

type homeRoutes struct{}

var homePattern = regexp.MustCompile(`^\/$`)

func (h homeRoutes) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch true {
	case r.Method == MethodGET && homePattern.MatchString(r.URL.Path):
		w.Write([]byte("GET /home"))
		return
	default:
		notFound(w, r)
	}
}

func LoadRoutes(server *http.ServeMux) {
	server.Handle("/", homeRoutes{})
	server.Handle("/users", usersRoute{})
	server.Handle("/users/", usersRoute{})
}
