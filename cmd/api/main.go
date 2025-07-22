package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"

	"github.com/dv1704/url_short/internal/db"
	"github.com/dv1704/url_short/internal/router"
)

func main() {
	// Load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("❌ Failed to load .env file: %v", err)
	}

	// 1. Initialize the database connection.
	db.InitDB()

	// 2. Get the DB object using the new GetDB() function.
	database := db.GetDB()

	// 3. Defer the closing of the database connection.
	sqlDB, err := database.DB()
	if err != nil {
		log.Fatalf("❌ Failed to get sql.DB: %v", err)
	}
	defer sqlDB.Close()

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	// This is the middleware that passes the database to the handlers
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", database)
		return c.Next()
	})

	router.SetupRoutes(app)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = ":3000"
	}

	log.Printf("🚀 Server running on %s", port)
	log.Fatal(app.Listen(port))
}
