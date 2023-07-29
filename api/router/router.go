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
	books.Patch("/:id", SessionMiddleware(), handlers.UpdateBook)
	books.Post("/", SessionMiddleware(), handlers.CreateBook)
	books.Delete("/:id", SessionMiddleware(), handlers.DeleteBook)

	categories := api.Group("/categories")
	categories.Get("/", handlers.GetCategories)
	categories.Get("/:id", handlers.GetCategory)
	categories.Patch("/:id", SessionMiddleware(), handlers.UpdateCategory)
	categories.Post("/", SessionMiddleware(), handlers.CreateCategory)
	categories.Delete("/:id", SessionMiddleware(), handlers.DeleteCategory)

	authors := api.Group("/authors")
	authors.Get("/", handlers.GetAuthors)
	authors.Post("/", SessionMiddleware(), handlers.CreateAuthor)

	users := api.Group("/users")
	users.Post("/login", handlers.Login)
	users.Post("/register", handlers.Register)
	users.Get("/profile", SessionMiddleware(), handlers.Profile)
}
