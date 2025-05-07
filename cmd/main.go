package main

import (
	"fmt"
	"go-clean-arch-fiber-boilerplate/internal/app/handlers"
	"go-clean-arch-fiber-boilerplate/internal/app/repositories"
	"go-clean-arch-fiber-boilerplate/internal/app/services"
	"go-clean-arch-fiber-boilerplate/pkg/config"
	"go-clean-arch-fiber-boilerplate/pkg/database"

	"github.com/gofiber/fiber/v2"
)

func main() {

	config := config.NewEnvConfig()
	db := database.Init(config, database.DBMigrator)

	app := fiber.New()

	// Route Group
	server := app.Group("/api")

	// User
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	handlers.NewUserHandler(server.Group("/users"), userService)

	app.Listen(fmt.Sprintf(":%s", config.AppPort))
}
