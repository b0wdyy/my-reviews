package main

import (
	"api/initializers"
	"api/router"
	"log"

	"github.com/gofiber/fiber/v2"
)

func init() {
	initializers.LoadEnv()
	initializers.InitDB()
	initializers.InitAWSConfig()
}

func main() {
	app := fiber.New()

	router.SetupRouter(app)

	log.Fatal(app.Listen(":8080"))
}
