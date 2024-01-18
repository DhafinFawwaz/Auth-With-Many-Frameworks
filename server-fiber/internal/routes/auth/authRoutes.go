package authRoutes

import (
	authHandler "fiber-auth-template/internal/handlers/auth"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthPublicRoutes(router fiber.Router) {
	auth := router.Group("/auth")

	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)

	auth.Get("/login", authHandler.GetAllMahasiswa) // Debug
}
func SetupAuthProtectedRoutes(router fiber.Router) {
	auth := router.Group("/auth")
	auth.Post("/", authHandler.Authenticate)
}
