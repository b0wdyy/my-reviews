package models

import "time"

type Category struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	Description string    `json:"description"`
	Books       []*Book   `gorm:"many2many:books_categories;"`
	CreatedAt   time.Time `json:"created_at"`
}
