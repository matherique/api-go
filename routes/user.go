package routes

import (
	"net/http"
	"regexp"
)

type usersRoute struct{}

var urlpattern = map[string]*regexp.Regexp{
	"/": regexp.MustCompile(`^\/users[\/]*$`),
}

func (u usersRoute) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch true {
	case r.Method == MethodGET && urlpattern["/"].MatchString(r.URL.Path):
		u.List(w, r)
		return
	}
}

func (u usersRoute) List(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GET /user -> lista todos usuarios"))
}
