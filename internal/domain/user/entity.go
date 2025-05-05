package user

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	FirstName string    `gorm:"size:255"`
	LastName  string    `gorm:"size:255"`
	Email     string    `gorm:"size:255;uniqueIndex"`
	Password  string    `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
