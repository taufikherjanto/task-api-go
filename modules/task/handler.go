package task

import (
	"section_10_go_repository_service/helper"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type taskHandler struct {
	service Service
}

func NewTaskHandler(service Service) *taskHandler {
	return &taskHandler{service}
}

func (h *taskHandler) CreateTask(c *fiber.Ctx) error {
	var task Task
	if err := c.BodyParser(&task); err != nil {
		response := helper.ApiResponse("Failed to parse request body", fiber.StatusBadRequest, "error", nil)
		return c.Status(response.Meta.Code).JSON(response)
	}

	if err := h.service.CreateTask(&task); err != nil {
		response := helper.ApiResponse("Failed to create task", fiber.StatusInternalServerError, "error", nil)
		return c.Status(response.Meta.Code).JSON(response)
	}

	response := helper.ApiResponse("Task created successfully", fiber.StatusOK, "success", task)
	return c.Status(response.Meta.Code).JSON(response)
}

func (h *taskHandler) GetAllTasks(c *fiber.Ctx) error {
	tasks, err := h.service.GetAllTasks()
	if err != nil {
		response := helper.ApiResponse("Failed to get tasks", fiber.StatusInternalServerError, "error", nil)
		return c.Status(response.Meta.Code).JSON(response)
	}

	response := helper.ApiResponse("Tasks fet successfully", fiber.StatusOK, "success", tasks)
	return c.Status(response.Meta.Code).JSON(response)

}

func (h *taskHandler) GetTaskByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		response := helper.ApiResponse("Invalid task ID", fiber.StatusBadRequest, "error", nil)
		return c.Status(response.Meta.Code).JSON(response)
	}

	task, err := h.service.GetTaskByID(id)
	if err != nil {
		response := helper.ApiResponse("Failed to fetch task", fiber.StatusNotFound, "error", nil)
		return c.Status(response.Meta.Code).JSON(response)
	}

	response := helper.ApiResponse("Task found", fiber.StatusOK, "success", task)
	return c.Status(response.Meta.Code).JSON(response)
}

func (h *taskHandler) UpdateTask(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		response := helper.ApiResponse("Invalid task ID", fiber.StatusBadRequest, "error", nil)
		return c.Status(response.Meta.Code).JSON(response)
	}

	var updateTask Task
	if err := c.BodyParser(&updateTask); err != nil {
		response := helper.ApiResponse("Failed to update task", fiber.StatusBadRequest, "error", nil)
		return c.Status(response.Meta.Code).JSON(response)
	}

	updateTask.ID = id
	if err := h.service.UpdateTask(&updateTask); err != nil {
		response := helper.ApiResponse("Failed to update task", fiber.StatusInternalServerError, "error", nil)
		return c.Status(response.Meta.Code).JSON(response)
	}

	response := helper.ApiResponse("Task updated successfully", fiber.StatusOK, "success", nil)
	return c.Status(response.Meta.Code).JSON(response)
}

func (h *taskHandler) DeleteTask(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		response := helper.ApiResponse("Invalid task ID", fiber.StatusBadRequest, "error", nil)
		return c.Status(response.Meta.Code).JSON(response)
	}

	if err := h.service.DeleteTask(id); err != nil {
		response := helper.ApiResponse("Failed to delete task", fiber.StatusInternalServerError, "error", nil)
		return c.Status(response.Meta.Code).JSON(response)
	}

	response := helper.ApiResponse("Task deleted successfully", fiber.StatusOK, "success", nil)
	return c.Status(response.Meta.Code).JSON(response)
}
