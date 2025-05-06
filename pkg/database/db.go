package database

import (
	"fmt"
	"go-clean-arch-fiber-boilerplate/pkg/config"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init(config *config.EnvConfig, DBMigrator func(db *gorm.DB) error) *gorm.DB {

	// dsn := fmt.Sprintf(`host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Shanghai`)
	dsn := fmt.Sprintf(
		`%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local`,
		config.DBUsername, config.DBPassword, config.DBHost, config.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		// DryRun: false,
	})

	if err != nil {
		log.Fatalf("Connect Database Fail !!!: %e", err)
	}

	log.Info("Connected to the database...")

	if err := DBMigrator(db); err != nil {
		log.Fatalf("Unable to migrate tables !!!: %e", err)
	}

	return db
}
