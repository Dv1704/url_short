package handler

import (
	"net/http"

	"github.com/dv1704/url_short/internal/db"
	"github.com/dv1704/url_short/internal/model"
	"github.com/gofiber/fiber/v2"
)

// ResolveURL handles redirection from short URL to original
func ResolveURL(c *fiber.Ctx) error {
	short := c.Params("url")

	var url model.URL
	result := db.DB.Where("short_url = ?", short).First(&url)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"error": "Short URL not found",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database query failed",
		})
	}

	// Increment redirect count
	url.RedirectCount++
	db.DB.Save(&url)

	return c.Redirect(url.OriginalURL, http.StatusMovedPermanently) // 301
}
