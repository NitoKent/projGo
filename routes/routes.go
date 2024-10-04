package routes

import (
	"example.com/m/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	app.Post("/books", handlers.CreateBook)
	app.Get("/books", handlers.GetBooks)
	app.Get("/books/:id", handlers.GetBookById)
	app.Put("/books/:id", handlers.UpdateBook)
	app.Post("/books/:id", handlers.RemoveBook)
}
