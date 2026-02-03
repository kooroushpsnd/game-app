package entity

type Item struct {
	ID uint `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"not null"`
	Description string `gorm:"not null"`
	Status bool `gorm:"not null;default:true"`
	UserId uint `gorm:"not null"`
}