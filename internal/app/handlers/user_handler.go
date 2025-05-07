package handlers

import (
	"go-clean-arch-fiber-boilerplate/internal/app/models"
	"go-clean-arch-fiber-boilerplate/internal/app/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	CreateUser(c *fiber.Ctx) error
	GetUserByID(c *fiber.Ctx) error
	ListUsers(c *fiber.Ctx) error
}

type userHandler struct {
	service services.UserService
}

func NewUserHandler(router fiber.Router, userService services.UserService) UserHandler {
	handler := &userHandler{
		service: userService,
	}

	router.Get("/:id", handler.GetUserByID)
	router.Post("/", handler.CreateUser)
	router.Get("/", handler.ListUsers)

	return handler
}

func (h *userHandler) CreateUser(c *fiber.Ctx) error {

	var input models.UserInputValidate

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "format invalid",
			"error":   err.Error(),
		})
	}

	user, err := h.service.CreateUser(input.Name, input.Email, input.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

func (h *userHandler) GetUserByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("invalid user id")
	}

	user, err := h.service.GetUserByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("user not found")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

func (h *userHandler) ListUsers(c *fiber.Ctx) error {
	users, err := h.service.GetUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    users,
	})
}
