package database

import (
	"go-clean-arch-fiber-boilerplate/internal/app/models"

	"gorm.io/gorm"
)

func DBMigrator(db *gorm.DB) error {
	return db.AutoMigrate(&models.User{})
}
