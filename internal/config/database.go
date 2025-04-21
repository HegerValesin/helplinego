package config

import (
	"fmt"
	"log"
	"os"

	"github.com/hegervalesin/helplinego/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupDatabase() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.Exec("SET CONSTRAINTS ALL DEFERRED")
	// err = db.AutoMigrate(&models.User{})
	// if err != nil {
	// 	log.Fatalf("Failed to migrate users table: %v", err)
	// }

	// Automigrate para criar as tabelas
	err = db.AutoMigrate(
		&models.Sector{},
		&models.User{},
		&models.Services{},
		&models.Facility{},
		&models.Floor{},
		&models.Room{},
	)

	db.Exec("SET CONSTRAINTS ALL IMMEDIATE")

	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	return db
}
