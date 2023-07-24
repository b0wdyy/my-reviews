package handlers

import (
	"api/services"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	return services.Login(c)
}

func Register(c *fiber.Ctx) error {
	return services.Register(c)
}

func Profile(c *fiber.Ctx) error {
	return c.SendString("Profile")
}
