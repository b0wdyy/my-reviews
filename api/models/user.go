package models

import "time"

type User struct {
	ID        string    `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `gorm:"not null" json:"password"`
	Name      string    `gorm:"not null" json:"name"`
	CreatedAt time.Time `gorm:"default:now()" json:"created_at"`
	DeletedAt time.Time `gorm:"default:null" json:"deleted_at"`
}
