package controller

import (
	"sigma-test/internal/request"
	"sigma-test/internal/response"
)

// Easely extensible with other services
type Services struct {
	UserService UserService
}

// Interface of methods for working with user service
type UserService interface {
	GetAllUsers() ([]response.User, error)
	GetUserById(id string) (response.User, error)
	CreateUser(user request.User) (response.User, error)
	UpdateUser(id string, user request.User) (response.User, error)
	DeleteUser(email string) error
}
