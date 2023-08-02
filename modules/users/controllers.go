package users

import "github.com/gofiber/fiber/v2"

func CreateUser(c *fiber.Ctx) error {

	payload := struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	return c.JSON(&payload)
}
