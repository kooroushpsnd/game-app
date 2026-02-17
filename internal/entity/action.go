package entity

import "time"

type Action struct {
	ID     uint 
	ItemID uint
	Date   time.Time 
	UserID uint
	Status ActionStatus
}

type ActionStatus string

const (
	ActionStatusCompleted ActionStatus = "completed"
	ActionStatusFailed    ActionStatus = "failed"
)