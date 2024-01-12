package service

import "sigma-test/internal/entity"

type Storages struct {
	UserRepo UserRepo
}

// Interface of methods for working with database
type UserRepo interface {
	AddUser(user entity.User) (entity.User, error)
	DeleteUser(idx int) error
	GetUserByIdx(idx int) (entity.User, error)
	GetAllUsers() ([]entity.User, error)
	UpdateUser(user entity.User, idx int) (entity.User, error)
	GetUserIdxById(id string) (int, bool)
	GetUserIdxByEmail(email string) (int, bool)
}
