package authHandler

import (
	"fiber-auth-template/database"
	"time"

	"github.com/gofiber/fiber/v2"
)

type mahasiswa struct {
	ID           int       `json:"id"`
	Password     string    `json:"password"`
	Last_login   time.Time `json:"last_login"`
	Is_superuser bool      `json:"is_superuser"`
	Username     string    `json:"username"`
	First_name   string    `json:"first_name"`
	Last_name    string    `json:"last_name"`
	Email        string    `json:"email"`
	Is_staff     bool      `json:"is_staff"`
	Is_active    bool      `json:"is_active"`
	Date_joined  time.Time `json:"date_joined"`
	Nim          string    `json:"nim"`
}

func GetAllMahasiswa(c *fiber.Ctx) error {
	data, _, _ := database.DB.From("user_api_mahasiswa").Select("*", "exact", false).Execute()
	err := c.Send(data)
	return err
}
func Register(c *fiber.Ctx) error {
	payload := mahasiswa{}
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	data, _, _ := database.DB.From("user_api_mahasiswa").Insert(map[string]interface{}{
		"id":           payload.ID,
		"password":     payload.Password,
		"last_login":   payload.Last_login,
		"is_superuser": payload.Is_superuser,
		"username":     payload.Username,
		"first_name":   payload.First_name,
		"last_name":    payload.Last_name,
		"email":        payload.Email,
		"is_staff":     payload.Is_staff,
		"is_active":    payload.Is_active,
		"date_joined":  payload.Date_joined,
		"nim":          payload.Nim,
	}, true, "", "", "",
	).Execute()

	c.Send(data)
	return nil
}
func Login(c *fiber.Ctx) error {
	payload := struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}{}
	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	var res []mahasiswa
	count, _ := database.DB.From("user_api_mahasiswa").Select("*", "exact", false).Filter("email", "=", payload.Email).ExecuteTo(&res)

	// case email not found
	if count == 0 {
		c.SendString("Email not found")
		return nil
	}

	// case wrong password
	if res[0].Password != payload.Password {
		c.SendString("Wrong password")
	}

	// c.SendString(string(res[0]))
	c.JSON(fiber.Map{
		"id":           res[0].ID,
		"password":     res[0].Password,
		"last_login":   res[0].Last_login,
		"is_superuser": res[0].Is_superuser,
		"username":     res[0].Username,
		"first_name":   res[0].First_name,
		"last_name":    res[0].Last_name,
		"email":        res[0].Email,
		"is_staff":     res[0].Is_staff,
		"is_active":    res[0].Is_active,
		"date_joined":  res[0].Date_joined,
		"nim":          res[0].Nim,
	})
	return nil
}
func Authenticate(c *fiber.Ctx) error {
	return nil
}
