package initializers

import (
	"log"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "postgres://postgres:jPbQQ6bxxQg3H4wSZEZq@db.yjaqrxnqrfgskgtkhkkz.supabase.co:5432/postgres?sslmode=disable"
	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	DB.Callback().Create().Before("gorm:before_create").Register("setID", setIDBeforeCreate)
}

func setIDBeforeCreate(tx *gorm.DB) {
	tx.Statement.SetColumn("ID", uuid.New().String())
}
