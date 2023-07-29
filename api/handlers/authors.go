package handlers

import (
	"api/services"

	"github.com/gofiber/fiber/v2"
)

func GetAuthors(c *fiber.Ctx) error {
	authors := services.GetAuthors()

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"status": "success",
		"data":   authors,
	})
}

func CreateAuthor(c *fiber.Ctx) error {
	author, err := services.CreateAuthor(c)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   author,
	})
}
