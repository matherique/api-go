package routes

import (
	"net/http"
	"regexp"

	"github.com/matherique/api-go/controllers"
)

type usersRoute struct {
	controller controllers.UserController
}

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
	users, err := u.controller.Index()
	if err != nil {
		internalServerError(w, r)
	}

	ok(w, r, users)
}
