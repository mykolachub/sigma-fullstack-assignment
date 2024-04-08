package controller

import (
	"sigma-test/config"
	"sigma-test/internal/request"
	"sigma-test/internal/response"

	"github.com/adrianbrad/queue"
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
	SignUp(body request.User) (response.User, config.ServiceCode, error)
	Login(body request.User) (string, config.ServiceCode, error)
	GetAllUsers(page int, search string) ([]response.User, config.ServiceCode, error)
	GetUserById(id string) (response.User, config.ServiceCode, error)
	GetUserByEmail(email string) (response.User, config.ServiceCode, error)
	CreateUser(user request.User) (response.User, config.ServiceCode, error)
	UpdateUser(id string, user request.User) (response.User, config.ServiceCode, error)
	DeleteUser(id string) (response.User, config.ServiceCode, error)
}

type PageService interface {
	TrackPage(q *queue.Linked[string], name string) (config.ServiceCode, error)
	GetPageCount(name string) (response.Page, config.ServiceCode, error)
}
