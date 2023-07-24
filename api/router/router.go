package router

import (
	"api/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {
	api := app.Group("/api")

	books := api.Group("/books")
	books.Get("/", handlers.GetBooks)
	books.Get("/:id", handlers.GetBook)
	books.Patch("/:id", handlers.UpdateBook)
	books.Post("/", handlers.NewBook)
	books.Delete("/:id", handlers.DeleteBook)

	categories := api.Group("/categories")
	categories.Get("/", handlers.GetCategories)
	categories.Get("/:id", handlers.GetCategory)
	categories.Patch("/:id", handlers.UpdateCategory)
	categories.Post("/", handlers.NewCategory)
	categories.Delete("/:id", handlers.DeleteCategory)

	users := api.Group("/users")
	users.Post("/login", handlers.Login)
	users.Post("/register", handlers.Register)
	users.Get("/profile", handlers.Profile)
}
