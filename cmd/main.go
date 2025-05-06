package main

import (
	"fmt"
	"go-clean-arch-fiber-boilerplate/pkg/config"
	"go-clean-arch-fiber-boilerplate/pkg/database"

	"github.com/gofiber/fiber/v2"
)

func main() {

	config := config.NewEnvConfig()
	db := database.Init(config, database.DBMigrator)
	_ = db

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("fiber hexagonal")
	})

	app.Listen(fmt.Sprintf(":%s", config.AppPort))
}
