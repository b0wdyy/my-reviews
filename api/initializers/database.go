package initializers

import (
	"log"
	"os"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	InitEnv()
}

func InitDB() {
	var (
		host     = os.Getenv("DB_HOST")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		dbname   = os.Getenv("DB_NAME")
		port     = os.Getenv("DB_PORT")
	)
	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port

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
