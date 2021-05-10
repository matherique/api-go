package controllers

import (
	"github.com/matherique/api-go/domain"
	"github.com/matherique/api-go/repository"
)

type UserController struct {
	Repository repository.UserRepository
}

func (controller UserController) Index() ([]domain.User, error) {
	return controller.Repository.FindAll()
}
