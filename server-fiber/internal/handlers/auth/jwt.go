package authHandler

import (
	"encoding/json"
	"fiber-auth-template/config"
	"fiber-auth-template/internal/models"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(mahasiswa models.Mahasiswa) (string, error) {
	// Create the Claims
	claims := jwt.MapClaims{
		"id":              mahasiswa.ID,
		"password":        mahasiswa.Password,
		"username":        mahasiswa.Username,
		"email":           mahasiswa.Email,
		"nim":             mahasiswa.NIM,
		"expiration_date": time.Now().Add(time.Hour * 24 * 100).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.GetEnv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return t, nil
}

func DecodeJWT(c *fiber.Ctx) (models.Mahasiswa, error) {

	// Get user from JWT middleware
	u := c.Locals("user")
	if u == nil {
		fmt.Println("No token provided")
		return models.Mahasiswa{}, c.Status(401).SendString("No token provided")
	}

	user := u.(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	// map[string]interface{} -> JSON
	jsonData, err := json.Marshal(claims)
	if err != nil {
		fmt.Println("Error marshalling claims\n", err)
		return models.Mahasiswa{}, err
	}

	// JSON -> struct
	mahasiswa := models.Mahasiswa{}
	err = json.Unmarshal(jsonData, &mahasiswa)
	if err != nil {
		fmt.Println("Error unmarshalling claims\n", err)
		return models.Mahasiswa{}, err
	}

	return mahasiswa, nil
}
