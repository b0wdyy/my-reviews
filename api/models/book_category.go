package models

import "time"

type BookCategory struct {
	BookID     string    `json:"book_id"`
	CategoryID string    `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
}
