package entity

import "time"

type Transaction struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Name      string
	Amount    float64
	UserId    uint
	CreatedAt time.Time
}
