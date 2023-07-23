package models

import (
	"time"
)

type Book struct {
	ID         string     `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Title      string     `json:"title"`
	Pages      int        `json:"pages"`
	Notes      string     `json:"notes"`
	Finished   bool       `json:"finished"`
	CoverImage string     `json:"cover_image"`
	Categories []Category `json:"categories" gorm:"many2many:books_categories;"`
	CreatedAt  time.Time  `json:"created_at"`
	DeletedAt  time.Time  `json:"deleted_at" gorm:"default:null"`
}
