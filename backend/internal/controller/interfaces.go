package controller

import (
	"sigma-test/internal/request"
	"sigma-test/internal/response"
)

// Easely extensible with other services
type Services struct {
	UserService UserService
}

type Configs struct {
	UserHandlerConfig UserHandlerConfig
}

// Interface of methods for working with user service
//
//go:generate mockery --name=UserService
type UserService interface {
	SignUp(body request.User) (response.User, error)
	Login(body request.User) (string, error)
	GetAllUsers() ([]response.User, error)
	GetUserById(id string) (response.User, error)
	GetUserByEmail(email string) (response.User, error)
	CreateUser(user request.User) (response.User, error)
	UpdateUser(id string, user request.User) (response.User, error)
	DeleteUser(id string) (response.User, error)
}
