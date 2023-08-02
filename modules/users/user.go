package users

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Init(app *fiber.App, DB *gorm.DB) {

	// run database migrations
	Migrate(DB)

	// init routes
	Register(app)

}
