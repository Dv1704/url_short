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

	_ "github.com/dv1704/url_short/docs" // 👈 Import generated Swagger docs
	"github.com/gofiber/swagger"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("⚠️  .env not found: %v", err)
	}

	// Initialize database
	db.InitDB()
	database := db.GetDB()
	sqlDB, err := database.DB()
	if err != nil {
		log.Fatalf("❌ Failed to get sql.DB: %v", err)
	}
	defer sqlDB.Close()

	// Initialize Fiber app
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())

	// Attach DB instance to context
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", database)
		return c.Next()
	})

	// Setup API routes
	router.SetupRoutes(app)

	// ✅ Serve Swagger JSON/YAML first
	app.Static("/docs/swagger.json", "./docs/swagger.json")
	app.Static("/docs/swagger.yaml", "./docs/swagger.yaml")

	// ✅ Serve Swagger UI after static routes
	app.Get("/docs/*", swagger.New(swagger.Config{
		URL: "/docs/swagger.json", // Path to the Swagger JSON
	}))

	// Get port from environment or default to :3000
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = ":3000"
	}

	log.Printf("🚀 Server running on %s", port)
	log.Fatal(app.Listen(port))
}
