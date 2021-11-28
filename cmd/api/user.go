package api

import (
	"net/http"
	"regexp"
)

type usersRoute struct{}

var urlpattern = map[string]*regexp.Regexp{
	"/": regexp.MustCompile(`^/\api\/users[\/]*$`),
}

func (u usersRoute) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch true {
	case r.Method == MethodGET && urlpattern["/"].MatchString(r.URL.Path):
		u.list(w, r)
		return
	default:
		notFound(w, r)
		return
	}
}

func (u usersRoute) list(w http.ResponseWriter, r *http.Request) {
	// tratar dados

	// mandar para /internal/app

	// pegar resposta e devolver
	ok(w, r, nil)
}
