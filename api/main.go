package main

import (
	"api/initializers"
	"api/router"
	"log"

	"github.com/gofiber/fiber/v2"
)

func init() {
	initializers.InitDB()
}

func main() {
	app := fiber.New()

	router.SetupRouter(app)

	log.Fatal(app.Listen(":8080"))
}
