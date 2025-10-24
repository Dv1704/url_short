package handler

import (
	"net/http"

	"github.com/dv1704/url_short/internal/db"
	"github.com/dv1704/url_short/internal/model"
	"github.com/gofiber/fiber/v2"
)

// ResolveURL handles redirection from short code to original URL
func ResolveURL(c *fiber.Ctx) error {
	short := c.Params("url")

	var url model.URL
	result := db.DB.Where("short_code = ?", short).First(&url)
	if result.Error != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"error": "Short URL not found",
		})
	}

	// Increment redirect count (optional)
	url.RedirectCount++
	db.DB.Save(&url)

	return c.Redirect(url.OriginalURL, http.StatusMovedPermanently) // 301 redirect
}
