package model

import (
	"time"
)

type User struct {
	ID        string
	Username  string
	Email     string
	Name      string
	Bio       string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
