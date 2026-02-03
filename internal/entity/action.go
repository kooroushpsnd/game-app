package entity

import "time"

type Action struct {
	ID     uint `gorm:"primaryKey;autoIncrement"`
	ItemID uint `gorm:"not null"`
	Date   time.Time `gorm:"not null"`
	UserID uint `gorm:"not null"`
	Status ActionStatus `gorm:"not null"`
}

type ActionStatus string

const (
	ActionStatusCompleted ActionStatus = "completed"
	ActionStatusFailed    ActionStatus = "failed"
)