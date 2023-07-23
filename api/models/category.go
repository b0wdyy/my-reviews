package models

import (
	"time"
)

type Category struct {
	ID          string    `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Description string    `json:"description"`
	Books       []Book    `gorm:"many2many:books_categories;"`
	CreatedAt   time.Time `json:"created_at"`
	DeletedAt   time.Time `json:"deleted_at" gorm:"default:null"`
}
