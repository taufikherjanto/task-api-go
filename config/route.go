package config

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	task.TaskRoutes(app, db)

	port := os.Getenv("PORT")
	portStr := ":" + port
	app.Listen(portStr)
}
