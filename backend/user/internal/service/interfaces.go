package service

import (
	"sigma-user/internal/entity"

	"github.com/adrianbrad/queue"
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
	BatchTrackPages(q *queue.Linked[string]) error
}
