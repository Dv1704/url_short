package main

import (
	"log"
	"os"

	"github.com/dv1704/url_short/internal/db"
	"github.com/dv1704/url_short/internal/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed  to load .env %v", err)
	}
	app := fiber.New()
	app.Use(logger.New())
	db.InitDB()
	router.SetupRoutes(app)
	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
