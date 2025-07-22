package handler

import (
	"os"
	"strings"
	"time"

	valid "github.com/asaskevich/govalidator"
	"github.com/dv1704/url_short/internal/db"
	"github.com/dv1704/url_short/internal/model"
	"github.com/dv1704/url_short/internal/util"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type request struct {
	URL         string `json:"url"`
	CustomShort string `json:"short"`
	Expiry      int64  `json:"expiry"`  // In hours
	UserID      uint   `json:"user_id"` 
}

type response struct {
	URL             string        `json:"url"`
	ShortURL        string        `json:"short"`
	RedirectCount   int           `json:"redirect_count"`
	CreatedAt       time.Time     `json:"created_at"`
	XRateRemaining  int           `json:"rate_limit"`
	XRateLimitReset time.Duration `json:"rate_limit_reset"`
}

func ShortenURL(c *fiber.Ctx) error {
	body := new(request)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON body",
		})
	}

	if !valid.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid URL format",
		})
	}

	if !util.RemoveDomainError(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot shorten your own domain",
		})
	}

	body.URL = util.EnforceHTTP(body.URL)

	var shortCode string
	if body.CustomShort == "" {
		shortCode = uuid.New().String()[:6]
	} else {
		shortCode = strings.ToLower(body.CustomShort)
	}

	var existing model.URL
	if err := db.DB.Where("short_code = ?", shortCode).First(&existing).Error; err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Short code already in use",
		})
	}

	expiryTime := time.Now().Add(24 * time.Hour)
	if body.Expiry > 0 {
		expiryTime = time.Now().Add(time.Duration(body.Expiry) * time.Hour)
	}

	// ✅ Extract user_id from JWT
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64)) // use float64 -> uint conversion

	entry := model.URL{
		OriginalURL:   body.URL,
		ShortCode:     shortCode,
		ExpiresAt:     &expiryTime,
		UserID:        userID,
		RedirectCount: 0,
	}

	if err := db.DB.Create(&entry).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to store shortened URL",
		})
	}

	shortURL := os.Getenv("DOMAIN") + "/" + shortCode

	return c.Status(fiber.StatusOK).JSON(response{
		URL:             body.URL,
		ShortURL:        shortURL,
		RedirectCount:   0,
		CreatedAt:       entry.CreatedAt,
		XRateRemaining:  10,
		XRateLimitReset: 30 * time.Minute,
	})
}

func GetUserURLs(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64)) // 🔁 Convert float64 to uint

	var urls []model.URL
	if err := db.DB.Where("user_id = ?", userID).Find(&urls).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to fetch URLs",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"urls":    urls,
	})
}
