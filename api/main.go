package main

import (
	"api/initializers"
	"api/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func init() {
	initializers.InitEnv()
	initializers.InitDB()
	initializers.InitAWSConfig()
	initializers.InitSession()
}

func main() {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(helmet.New())

	router.SetupRouter(app)

	log.Fatal(app.Listen(":8080"))
}
