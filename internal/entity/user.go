package entity

import (
	"time"
)

type User struct {
	ID          uint
	PhoneNumber string
	Email       string
	Name        string
	Password    string
	Role        Role
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
