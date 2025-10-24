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

	"github.com/gofiber/swagger"
)

func main() {
	// Load env variables
	err := godotenv.Load()
	if err != nil {
		log.Printf("⚠️  .env not found: %v", err)
	}

	// Initialize DB
	db.InitDB()
	database := db.GetDB()
	sqlDB, err := database.DB()
	if err != nil {
		log.Fatalf("❌ Failed to get sql.DB: %v", err)
	}
	defer sqlDB.Close()

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", database)
		return c.Next()
	})

	// Setup API routes
	router.SetupRoutes(app)

	// Serve swagger.json from the docs directory
	app.Static("/docs/swagger.json", "./docs/swagger.json")

	// Serve Swagger UI using the correct JSON file
	app.Get("/docs/*", swagger.New(swagger.Config{
		URL: "/docs/swagger.json", // URL to the Swagger JSON
	}))

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = ":3000"
	}

	log.Printf("🚀 Server running on %s", port)
	log.Fatal(app.Listen(port))
}
