package handlers

import (
	"api/initializers"
	"api/services"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func init() {
	initializers.InitDB()
}

func GetBooks(c *fiber.Ctx) error {
	books := services.GetBooks(c)

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   books,
	})
}

func GetBook(c *fiber.Ctx) error {
	book := services.GetBook(c.Params("id"))

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   book,
	})
}

func NewBook(c *fiber.Ctx) error {
	book, err := services.NewBook(c)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   book,
	})
}

func UpdateBook(c *fiber.Ctx) error {
	book, err := services.UpdateBook(c)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   book,
	})
}

func DeleteBook(c *fiber.Ctx) error {
	services.DeleteBook(c.Params("id"))

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status": "success",
	})
}
