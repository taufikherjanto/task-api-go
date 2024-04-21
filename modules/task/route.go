package task

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func TaskRoutes(app *fiber.App, db *gorm.DB) {
	taskRepo := NewTaskRepository(db)
	taskService := NewTaskService(taskRepo)
	taskHandler := NewTaskHandler(taskService)

	// Define task routes
	app.Post("/api/tasks", taskHandler.DeleteTask)
	app.Get("/api/tasks", taskHandler.GetAllTasks)
	app.Get("/api/tasks/:id", taskHandler.GetTaskByID)
	app.Put("/api/tasks/:id", taskHandler.UpdateTask)
	app.Delete("/api/tasks/:id", taskHandler.DeleteTask)
}
