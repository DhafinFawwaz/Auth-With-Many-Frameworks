package router

import (
	authRoutes "fiber-auth-template/internal/routes/auth"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
	authRoutes.SetupAuthRoutes(api)
}
