package handlers

import (
	"api/services"

	"github.com/gofiber/fiber/v2"
)

func GetCategories(c *fiber.Ctx) error {
	return c.SendString("All categories")
}

func GetCategory(c *fiber.Ctx) error {
	return c.SendString("One category")
}

func CreateCategory(c *fiber.Ctx) error {
	category, err := services.CreateCategory(c.FormValue("description"))

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Category already exists",
			"status": "error",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"data":   category,
		"status": "success",
	})
}

func UpdateCategory(ctx *fiber.Ctx) error {
	return ctx.SendString("Update category")
}

func DeleteCategory(ctx *fiber.Ctx) error {
	return ctx.SendString("Delete category")
}
