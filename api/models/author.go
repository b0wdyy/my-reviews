package models

import (
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name    string `json:"name"`
	Books   []Book `json:"books"`
	Picture string `json:"picture"`
}
