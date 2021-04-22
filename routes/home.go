package routes

import (
	"log"
	"net/http"
	"regexp"
)

type homeRoutes struct {
	logger *log.Logger
}

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

func NewHomeRoute(logger *log.Logger) usersRoute {
	return usersRoute{
		logger: logger,
	}
}
