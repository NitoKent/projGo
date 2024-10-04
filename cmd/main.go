package main

import (
	"example.com/m/config"
	"example.com/m/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	config.InitDatabase()
	app := fiber.New()
	routes.SetupRoutes(app)
	app.Listen(":3010")
}
