package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/dv1704/url_short/internal/model"
)

var DB *gorm.DB

func InitDB() {
	dsn := os.Getenv("DATABASE_URL")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	// Ping test
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get sql.DB from gorm.DB: %v", err)
	}
	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Failed to ping DB: %v", err)
	}

	
	err = DB.AutoMigrate(
		&model.User{},
		&model.URL{},

	)
	if err != nil {
		log.Fatalf("Failed to auto-migrate: %v", err)
	}

	fmt.Println("Successfully connected and migrated the DB")
}
