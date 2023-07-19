package main

import (
	"api/router"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	app := fiber.New()

	router.SetupRouter(app)

	log.Fatal(app.Listen(":3000"))
}
