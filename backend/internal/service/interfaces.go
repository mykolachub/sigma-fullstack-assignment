package service

import (
	"sigma-test/internal/entity"
)

type Storages struct {
	UserRepo UserRepo
	PageRepo PageRepo
}

// Interface of methods for working with database
type UserRepo interface {
	CreateUser(user entity.User) (entity.User, error)
	GetUsers(page int, search string) ([]entity.User, error)
	GetUser(id string) (entity.User, error)
	GetUserByEmail(email string) (entity.User, error)
	UpdateUser(id string, data entity.User) (entity.User, error)
	DeleteUser(id string) (entity.User, error)
}

type PageRepo interface {
	GetPage(name string) (entity.Page, error)
	ResetPageCount(name string) error
	TrackPage(name string) error
}
