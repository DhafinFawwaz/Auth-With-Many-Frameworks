package main

import (
	"fiber-auth-template/config"
	"fiber-auth-template/database"
	"fiber-auth-template/router"
	"os"
	"os/signal"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:      "Auth Template",
		ServerHeader: "Fiber",
	})

	database.ConnectDatabase()
	defer database.DisconnectDatabase()

	app.Use(logger.New())

	router.SetupPublicRoutes(app)
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(config.GetEnv("JWT_SECRET"))},
	}))
	router.SetupProtectedRoutes(app)

	// Listen to exit
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)
		<-quit
		database.DisconnectDatabase()
		os.Exit(0)
	}()

	app.Listen(":8080")
}
