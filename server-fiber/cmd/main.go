package main
import (
    "net/http"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
)
func main() {
    app := fiber.New(fiber.Config{
        AppName:      "Fiber with Planetscale",
        ServerHeader: "Fiber",
    })
    app.Use(logger.New())
    app.Get("/", func(c *fiber.Ctx) error {
        return c.Status(http.StatusOK).JSON(&fiber.Map{
            "message": "Hello world",
        })
    })
    app.Listen(":3000")
}