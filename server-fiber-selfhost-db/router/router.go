package router

import (
	debugHandler "fiber-auth-template/internal/handlers/debug"
	authRoutes "fiber-auth-template/internal/routes/auth"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupPublicRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
	authRoutes.SetupAuthPublicRoutes(api)

	api.Get("/sql", debugHandler.SQL) // Debug
}

func SetupProtectedRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())
	authRoutes.SetupAuthProtectedRoutes(api)
}
