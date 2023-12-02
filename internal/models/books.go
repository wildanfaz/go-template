package models

import "time"

type (
	Book struct {
		Id          uint      `json:"id"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Author      string    `json:"author"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}

	Books []Book
)
