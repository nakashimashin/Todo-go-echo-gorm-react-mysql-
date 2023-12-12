package model

import "time"

type Task struct {
	ID        uint      `json:"id" param:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
