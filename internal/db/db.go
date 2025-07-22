package db

import (
    "fmt"
    "log"
    "os"

    "github.com/dv1704/url_short/internal/model"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
    dsn := os.Getenv("DATABASE_URL")
    var err error

    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("❌ Failed to connect to DB: %v", err)
    }

    sqlDB, err := DB.DB()
    if err != nil {
        log.Fatalf("❌ Failed to get sql.DB: %v", err)
    }
    if err := sqlDB.Ping(); err != nil {
        log.Fatalf("❌ Failed to ping DB: %v", err)
    }

    // AutoMigrate
    err = DB.AutoMigrate(&model.User{}, &model.URL{})
    if err != nil {
        log.Fatalf("❌ AutoMigrate failed: %v", err)
    }

    fmt.Println("✅ Database connected and migrated")
}


func GetDB() *gorm.DB {
    return DB
}