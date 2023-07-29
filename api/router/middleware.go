package router

import (
	"api/initializers"

	"github.com/gofiber/fiber/v2"
)

func init() {
	initializers.InitSession()
}

func SessionMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		session, err := initializers.Store.Get(c)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		if session.Get("user_id") == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		c.Locals("user_id", session.Get("user_id"))

		return c.Next()
	}
}
