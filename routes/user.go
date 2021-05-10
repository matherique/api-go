package routes

import (
	"database/sql"
	"log"
	"net/http"
	"regexp"

	"github.com/matherique/api-go/controllers"
	"github.com/matherique/api-go/repository"
)

type usersRoute struct {
	logger     *log.Logger
	controller controllers.UserController
}

var urlpattern = map[string]*regexp.Regexp{
	"/": regexp.MustCompile(`^\/users[\/]*$`),
}

func (u usersRoute) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	u.logger.Printf("%s - %s\n", r.Method, r.URL.Path)
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

func NewUserRoute(database *sql.DB, logger *log.Logger) usersRoute {
	repository := repository.UserRepository{
		Database: database,
	}

	controller := controllers.UserController{
		Repository: repository,
	}

	return usersRoute{
		controller: controller,
		logger:     logger,
	}
}
