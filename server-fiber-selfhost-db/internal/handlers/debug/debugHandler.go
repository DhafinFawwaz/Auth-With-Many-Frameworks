package debug

import (
	"fiber-auth-template/database"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func SQL(c *fiber.Ctx) error {
	query := c.Query("query")
	text, err := database.SQL(&query)
	if err != nil {
		fmt.Println("Error SQL Query\n", err)
		c.JSON(err)
		return err
	}
	c.JSON(text)
	return nil
}
