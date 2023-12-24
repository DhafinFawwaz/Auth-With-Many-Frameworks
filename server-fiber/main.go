package main

import (
	"fiber-auth-template/database"
	"fiber-auth-template/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:      "Auth Template",
		ServerHeader: "Fiber",
	})
	database.ConnectDatabase()
	app.Use(logger.New())
	router.SetupRoutes(app)
	app.Listen(":3000")
}
