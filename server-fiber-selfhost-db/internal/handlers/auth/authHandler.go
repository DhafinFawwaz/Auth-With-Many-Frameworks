package authHandler

import (
	"fiber-auth-template/database"
	"fiber-auth-template/internal/models"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetAllMahasiswa(c *fiber.Ctx) error {
	data, err := database.SelectAllMahasiswa()
	if err != nil {
		fmt.Println("Error getting all mahasiswa\n", err)
		return err
	}
	c.JSON(data)
	return nil
}
func Register(c *fiber.Ctx) error {
	newMahasiswa := models.Mahasiswa{}
	if err := c.BodyParser(&newMahasiswa); err != nil {
		fmt.Println("Error parsing body\n", err)
		return err
	}

	// check if email already exists
	_, err := database.SelectByEmail(newMahasiswa.Email)
	if err == nil {
		fmt.Println("Email already exists\n", err)
		return c.Status(409).SendString("Email already exists")
	}

	// Assign values to newMahasiswa
	newMahasiswa.DateJoined = time.Now()
	newMahasiswa.Password, err = EncryptPassword(newMahasiswa.Password)
	if err != nil {
		fmt.Println("Error encrypting password\n", err)
		return err
	}

	// Insert newMahasiswa to database
	err = database.InsertNewMahasiswa(newMahasiswa)
	if err != nil {
		fmt.Println("Error inserting mahasiswa\n", err)
		return err
	}

	c.JSON(newMahasiswa)
	return nil
}
func Login(c *fiber.Ctx) error {
	req := models.Mahasiswa{}
	if err := c.BodyParser(&req); err != nil {
		fmt.Println("Error parsing body", err)
		return err
	}

	// find email
	mahasiswa, err := database.SelectByEmail(req.Email)
	if err != nil {
		fmt.Println("Email not found\n", err)
		return c.Status(404).SendString("Email not found")
	}

	// Check encrypted password
	if !ComparePassword(req.Password, mahasiswa.Password) {
		fmt.Println("Wrong password\n", err)
		return c.Status(401).SendString("Wrong password")
	}

	// Generate JWT
	AccessToken, err := GenerateJWT(mahasiswa)
	if err != nil {
		fmt.Println("Error generating JWT\n", err)
		return err
	}
	authenticatedMahasiswa := models.AuthenticatedMahasiswa{
		Mahasiswa:   *mahasiswa,
		AccessToken: AccessToken,
	}

	c.JSON(authenticatedMahasiswa)
	return nil
}
func Authenticate(c *fiber.Ctx) error {
	mahasiswa, err := DecodeJWT(c)
	if err != nil {
		fmt.Println("Error decoding JWT\n", err)
		return err
	}
	c.JSON(mahasiswa)
	return nil
}
