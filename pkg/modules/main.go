package modules

import (
	"rismoon/modules/users"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Register(app *fiber.App, DB *gorm.DB) {

	// Users module
	users.Init(app, DB)
}
