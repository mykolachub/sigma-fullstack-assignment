package controller

import (
	"sigma-test/internal/request"
	"sigma-test/internal/response"
)

// Easely extensible with other services
type Services struct {
	UserService UserService
	PageService PageService
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
	GetAllUsers(page int, search string) ([]response.User, error)
	GetUserById(id string) (response.User, error)
	GetUserByEmail(email string) (response.User, error)
	CreateUser(user request.User) (response.User, error)
	UpdateUser(id string, user request.User) (response.User, error)
	DeleteUser(id string) (response.User, error)
}

type PageService interface {
	TrackPage(name string) error
	GetPageCount(name string) (response.Page, error)
}
