package model

import "time"

type Task struct {
	ID        uint      `json:"id" gorm:"primary_key" param:"id"`
	Title     string    `json:"title" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
