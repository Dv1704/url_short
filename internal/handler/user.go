package handler

import (
	"os"
	"time"

	"github.com/dv1704/url_short/internal/db"
	"github.com/dv1704/url_short/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// RegisterUser handles user signup
func RegisterUser(c *fiber.Ctx) error {
	db, ok := c.Locals("db").(*gorm.DB)
	if !ok || db == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "Database connection not available",
		})
	}

	var input model.User

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Cannot parse JSON",
		})
	}

	// Check if email already exists
	var existing model.User
	if err := db.Where("email = ?", input.Email).First(&existing).Error; err == nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"success": false,
			"error":   "Email already registered",
		})
	}

	// DO NOT manually hash here — GORM will hash it using BeforeCreate
	if err := db.Create(&input).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"error":   "Failed to create user",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "User registered successfully",
	})
}

func Login(c *fiber.Ctx) error {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Invalid login input",
		})
	}

	var user model.User
	if err := db.DB.Where("email = ?", body.Email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Invalid email or password",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Invalid email or password",
		})
	}

	// Create JWT
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(), // 3 days
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "Failed to generate token",
		})
	}

	// Set JWT in cookie
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    signedToken,
		Expires:  time.Now().Add(time.Hour * 72),
		HTTPOnly: true,
		Secure:   false, // set true in production
		Path:     "/",
	})

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Login successful",
	})
}
