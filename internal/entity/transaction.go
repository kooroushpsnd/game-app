package entity

import "time"

type Transaction struct {
	ID        uint
	Name      string
	Amount    float64
	UserId    uint
	CreatedAt time.Time
}
