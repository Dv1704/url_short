package router

import (
	"github.com/dv1704/url_short/internal/handler"
	"github.com/dv1704/url_short/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Authentication routes
	app.Post("/api/v0/signup", handler.RegisterUser)
	app.Post("/api/v0/login", handler.Login)

	// Public route for resolving short URLs
	app.Get("/:url", handler.ResolveURL)

	// Protected routes (require valid JWT)
	api := app.Group("/api/v0", middleware.JWTProtected())

	api.Post("/shorten", handler.ShortenURL) 
	api.Get("/my-urls", handler.GetUserURLs) 
}
