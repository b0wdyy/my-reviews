package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title      string     `json:"title"`
	Pages      int64      `json:"pages"`
	CoverImage string     `json:"cover_image"`
	Notes      string     `json:"notes"`
	Finished   bool       `json:"finished"`
	AuthorID   *uint      `json:"author_id"`
	Author     Author     `json:"author"`
	Categories []Category `json:"categories" gorm:"many2many:book_categories;"`
}
