package users

import (
	"rismoon/pkg/validate"

	"github.com/gofiber/fiber/v2"
)

type RegisterUserDTO struct {
	Email          string `validate:"required,email"`
	Password       string `validate:"required,min=5,max=32"`
	RepeatPassword string `validate:"required,min=5,max=32"`
}

func RegisterUser(c *fiber.Ctx) error {

	payload := struct {
		Email          string `json:"email"`
		Password       string `json:"password"`
		RepeatPassword string `json:"repeat_password"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	userData := RegisterUserDTO{
		Email:          payload.Email,
		Password:       payload.Password,
		RepeatPassword: payload.RepeatPassword,
	}

	passed, errs := validate.Check(userData)

	if !passed {
		return c.Status(400).JSON(errs)
	}

	return c.JSON(&payload)
}
