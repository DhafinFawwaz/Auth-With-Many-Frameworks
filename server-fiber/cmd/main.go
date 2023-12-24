package main
import (
    "net/http"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/maful/fiber-pscale/models"
)
func main() {
    app := fiber.New(fiber.Config{
        AppName:      "Fiber with Planetscale",
        ServerHeader: "Fiber",
    })
    
    models.ConnectDatabase()

    app.Use(logger.New())
    
    app.Listen(":3000")
}