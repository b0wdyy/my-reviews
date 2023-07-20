package handlers

import (
	"api/initializers"
	"api/models"

	"github.com/gofiber/fiber/v2"
)

func init() {
	initializers.InitDB()
}

func GetBooks(c *fiber.Ctx) error {
	var books []models.Book
	initializers.DB.Preload("Categories").Find(&books)

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   books,
	})
}

func GetBook(c *fiber.Ctx) error {
	return c.JSON("Single book")
}

func NewBook(c *fiber.Ctx) error {
	return c.JSON("New book")
}

func UpdateBook(c *fiber.Ctx) error {

	return c.JSON("Update book")
}

func DeleteBook(c *fiber.Ctx) error {
	return c.JSON("Delete book")
}
