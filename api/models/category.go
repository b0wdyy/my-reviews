package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Description string  `json:"description"`
	Books       []*Book `json:"books" gorm:"many2many:book_categories;"`
}
