package main

import (
	"log"
	"os"
	"rismoon/pkg/db"
	"rismoon/pkg/modules"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	// load env variables
	envError := godotenv.Load()
	if envError != nil {
		log.Fatal("Unable to read .env file")
	}

	// create fiber app instance
	app := fiber.New()

	// conbect to database
	Database := db.Register()

	// register modules
	modules.Register(app, Database)

	// listen app
	app.Listen(":" + os.Getenv("APP_PORT"))

}
