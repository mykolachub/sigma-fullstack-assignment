package service

import "sigma-test/internal/entity"

type Storages struct {
	UserRepo UserRepo
}

// Interface of methods for working with database
type UserRepo interface {
	CreateUser(user entity.User) (entity.User, error)
	GetUsers() ([]entity.User, error)
	GetUser(id string) (entity.User, error)
	GetUserByEmail(email string) (entity.User, error)
	UpdateUser(id string, data entity.User) (entity.User, error)
	DeleteUser(id string) (entity.User, error)
}
