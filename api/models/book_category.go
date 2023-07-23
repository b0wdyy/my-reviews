package models

import (
	"time"
)

type BookCategory struct {
	ID         string    `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	BookID     string    `json:"book_id"`
	CategoryID string    `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	DeletedAt  time.Time `json:"deleted_at" gorm:"default:null"`
}
