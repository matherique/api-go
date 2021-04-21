package controllers

import (
	"github.com/matherique/api-go/domain"
)

type UserController struct {
	repo interface{}
}

type User = domain.User

func (cntr UserController) Index() ([]User, error) {
	return []User{}, nil
}
