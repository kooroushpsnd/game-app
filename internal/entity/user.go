package entity

import (
	"time"
)

type User struct {
	ID          uint
	Email       string
	Name        string
	Password    string
	Role        Role
	Status		bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
