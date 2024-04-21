package main

import (
	"section_10_go_repository_service/config"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	app := fiber.New()

	// load environtment variables from .env file
	// use library godotenv
	// use godotenv.load()

	// connect to db
	db, err := config.ConnectDB()
	if err != nil {
		panic("Failed to connect to the database")
	}

	// setup routes
	config.SetupRoutes(app, db)
}
