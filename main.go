package main

import (
	"log"
	"os"
	"rismoon/pkg/modules"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Databaseconnection failed")
	}

	// register modules
	modules.Register(app, db)

	// listen app
	app.Listen(":3000")

}
