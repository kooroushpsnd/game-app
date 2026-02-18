package entity

import "time"

type EmailCode struct {
	ID             uint
	Email          string
	HashCode       string
	Status         bool
	Attempts       int
	ExpirationDate time.Time
	UserID         uint
	CreatedAt      time.Time
}
