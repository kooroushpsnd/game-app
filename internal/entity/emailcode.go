package entity

import "time"

type EmailCodeStatus string

const (
	EmailCodeStatusActive   EmailCodeStatus = "active"
	EmailCodeStatusVerified EmailCodeStatus = "verified"
	EmailCodeStatusExpired  EmailCodeStatus = "expired"
)

type EmailCode struct {
	ID             uint
	Email          string
	HashCode       string
	Status         EmailCodeStatus
	Attempts       int
	ExpirationDate time.Time
	UserID         uint
	CreatedAt      time.Time
}
