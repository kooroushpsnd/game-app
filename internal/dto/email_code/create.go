package emailcodedto

import "time"

type CreateEmailCodeDto struct {
	Email          string
	HashCode       string
	ExpirationDate time.Time
	UserID         uint
	Attempts	   int
}
