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

	// Import docs package (required for swag)
	_ "github.com/dv1704/url_short/docs"
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

	// 🔒 Only serve Swagger if swagger.json exists (avoids 404s & redirect loops)
	if _, err := os.Stat("./docs/swagger.json"); err == nil {
		app.Static("/docs/swagger.json", "./docs/swagger.json")
		app.Static("/docs/swagger.yaml", "./docs/swagger.yaml")
		app.Use("/docs", swagger.New(swagger.Config{
			URL: "/docs/swagger.json",
		}))
		log.Println("✅ Swagger UI enabled at /docs")
	} else {
		log.Println("⚠️  Swagger disabled: ./docs/swagger.json not found")
	}

	// Get port from environment (e.g., Render)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("🚀 Server running on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
