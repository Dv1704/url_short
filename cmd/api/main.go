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
	if err := godotenv.Load(); err != nil {
		log.Printf("⚠️  .env not found: %v", err)
	}

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

	router.SetupRoutes(app)

	// ✅ Serve ONLY swagger.json and Swagger UI
	if _, err := os.Stat("./docs/swagger.json"); err == nil {
		app.Static("/docs/swagger.json", "./docs/swagger.json")
		app.Use("/docs", swagger.New(swagger.Config{
			URL: "/docs/swagger.json",
		}))
		log.Println("✅ Swagger UI enabled at /docs")
	} else {
		log.Println("⚠️  Swagger disabled: ./docs/swagger.json not found")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("🚀 Server running on port %s", port)
	log.Fatal(app.Listen(":" + port))


	app.Get("/debug", func(c *fiber.Ctx) error {
    if _, err := os.Stat("docs/swagger.json"); err != nil {
        return c.SendString("MISSING")
    }
    return c.SendString("FOUND")
})
}
