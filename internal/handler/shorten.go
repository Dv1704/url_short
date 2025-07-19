package handler

import (
	"time"

	"github.com/gofiber/fiber/v2" 
)

type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

type response struct {
	URL             string        `json:"url"`
	CustomShort     string        `json:"short"`
	Expiry          time.Duration `json:"expiry"`
	XRateRemanining int           `json:"rate_limit"`
	XRateLimitReset time.Duration `json:"rate_limit_reset"`
}

func ShortenURL(c *fiber.Ctx) error {
	body := new(request)

	
	if err := c.BodyParser(body); err != nil {

		c.Status(fiber.StatusBadRequest) 
		return c.JSON(fiber.Map{        
			"error": "Invalid request body",
		})
	}

	resp := response{
		URL:             body.URL,
		CustomShort:     body.CustomShort,
		Expiry:          body.Expiry,
		XRateRemanining: 10,
		XRateLimitReset: 30 * time.Minute,
	}

	
	return c.Status(fiber.StatusOK).JSON(resp)
}