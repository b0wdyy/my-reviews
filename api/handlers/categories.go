package handlers

import "github.com/gofiber/fiber/v2"

func GetCategories(ctx *fiber.Ctx) error {
	return ctx.SendString("All categories")
}

func GetCategory(ctx *fiber.Ctx) error {
	return ctx.SendString("One category")
}

func NewCategory(ctx *fiber.Ctx) error {
	return ctx.SendString("New category")
}

func UpdateCategory(ctx *fiber.Ctx) error {
	return ctx.SendString("Update category")
}

func DeleteCategory(ctx *fiber.Ctx) error {
	return ctx.SendString("Delete category")
}
