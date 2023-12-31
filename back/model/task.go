package model

import "time"

type Task struct {
	ID        uint      `json:"id" gorm:"primaryKey" param:"id"`
	Title     string    `json:"title" gorm:"not null"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
