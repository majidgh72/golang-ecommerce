package users

import "github.com/gofiber/fiber/v2"

func Register(app *fiber.App) {

	// Create User
	app.Post("/users", CreateUser)
}
