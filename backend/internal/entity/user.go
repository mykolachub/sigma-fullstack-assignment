package entity

import (
	"sigma-test/internal/response"
	"time"
)

type User struct {
	ID        string
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
}

func (u *User) ToResponse() response.User {
	return response.User{
		ID:        u.ID,
		Email:     u.Email,
		Password:  u.Password,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
	}
}
