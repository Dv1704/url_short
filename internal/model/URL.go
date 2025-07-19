package model

import (
	"time"

	"gorm.io/gorm"
)
type URL struct {
    ID          uint           `gorm:"primaryKey"`
    ShortCode   string         `gorm:"uniqueIndex;not null"`
    OriginalURL string         `gorm:"not null"`
    UserID      uint           // foreign key
    CreatedAt   time.Time
    UpdatedAt   time.Time
    DeletedAt   gorm.DeletedAt `gorm:"index"`
}
