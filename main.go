package main

import (
	"rismoon/pkg/modules"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	// create fiber app instance
	app := fiber.New()

	// conbect to database
	dsn := "root:@tcp(127.0.0.1:3306)/go?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Databaseconnection failed")
	}

	// register modules
	modules.Register(app, db)

	// listen app
	app.Listen(":3000")

}
