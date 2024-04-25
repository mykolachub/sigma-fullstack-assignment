package entity

import (
	"time"
)

type User struct {
	ID        string
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
}
