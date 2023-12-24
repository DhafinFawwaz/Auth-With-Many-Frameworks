package authRoutes

import (
	authHandler "fiber-auth-template/internal/handlers/auth"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(router fiber.Router) {
	auth := router.Group("/auth")

	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)
	auth.Post("/", authHandler.Authenticate)

	auth.Get("/", authHandler.GetAllMahasiswa) // Debug
}
