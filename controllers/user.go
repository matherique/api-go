package controllers

import (
	"github.com/matherique/api-go/domain"
	"github.com/matherique/api-go/repository"
)

type UserController struct {
	Repository repository.UserRepository
}

type User = domain.User

func (controller UserController) Index() ([]User, error) {
	return controller.Repository.List()
}
