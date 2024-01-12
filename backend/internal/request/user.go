package request

import (
	"sigma-test/internal/entity"
)

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u User) ToEntity() entity.User {
	return entity.User{
		ID:       u.ID,
		Email:    u.Email,
		Password: u.Password,
	}
}
