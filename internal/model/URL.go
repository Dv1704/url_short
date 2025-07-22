package model

import (
	"time"

	"gorm.io/gorm"
)

type URL struct {
	ID            uint           `gorm:"primaryKey"`
	ShortCode     string         `gorm:"uniqueIndex;not null"`
	OriginalURL   string         `gorm:"not null"`
	UserID        uint           `gorm:"not null"`                       // Foreign key field
	User          User           `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` 
	RedirectCount int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
    ExpiresAt *time.Time `json:"expires_at"`

}
