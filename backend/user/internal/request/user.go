package request

import (
	"sigma-user/internal/entity"
)

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (u User) ToEntity() entity.User {
	return entity.User{
		ID:       u.ID,
		Email:    u.Email,
		Password: u.Password,
		Role:     u.Role,
	}
}
