package router

import (
	_ "github.com/dv1704/url_short/docs" // Swagger docs
	"github.com/dv1704/url_short/internal/handler"
	"github.com/dv1704/url_short/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger" // Swagger UI
)

func SetupRoutes(app *fiber.App) {
	// 🔹 Swagger Docs (must come first)
	app.Get("/swagger/*", swagger.HandlerDefault) // http://localhost:3000/swagger/index.html

	// 🔹 Authentication routes
	app.Post("/api/v0/signup", handler.RegisterUser)
	app.Post("/api/v0/login", handler.Login)

	// 🔹 Protected routes (require valid JWT)
	api := app.Group("/api/v0", middleware.JWTProtected())
	api.Post("/shorten", handler.ShortenURL)
	api.Get("/my-urls", handler.GetUserURLs)

	// 🔹 Public route for resolving short URLs (must come last!)
	// Only match 6-character alphanumeric short URLs to avoid conflicts
	app.Get("/:shortURL([a-zA-Z0-9]{6})", handler.ResolveURL)
}
