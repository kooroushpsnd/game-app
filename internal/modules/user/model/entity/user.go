package entity

import "time"

type User struct {
	ID          uint	`gorm:"primaryKey;autoIncrement"`
	PhoneNumber string	`gorm:"uniqueIndex;not null"`
	Email       string	`gorm:"uniqueIndex;not null"`
	Name        string	`gorm:"not null"`
	Password    string	`gorm:"not null"`
	Role        Role	`gorm:"not nulll;default:1"`
	CreatedAt   time.Time
	UpdatedAt   time.Time

	Transactions []Transaction `gorm:"foreignKey:UserId"`
}
