package router

import (
	"github.com/dv1704/url_short/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// app.Get("/:url", handler.ResolveURL)
	app.Post("/api/v0", handler.ShortenURL)
}
