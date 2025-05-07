package handlers

import (
	"go-clean-arch-fiber-boilerplate/internal/app/models"
	"go-clean-arch-fiber-boilerplate/internal/app/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(router fiber.Router, userService services.UserService) {
	handler := &UserHandler{
		service: userService,
	}

	router.Get("/:id", handler.GetUserByID)
	router.Post("/", handler.CreateUser)
	router.Get("/", handler.GetUserList)
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {

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
		"message": "user create successfully.",
		"data":    user,
	})
}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	user, err := h.service.GetUserByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString("user not found")
	}

	return c.JSON(user)
}

func (h *UserHandler) GetUserList(c *fiber.Ctx) error {
	users, err := h.service.GetUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(users)
}
