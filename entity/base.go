package entity

import (
	"time"

	"gorm.io/gorm"
)

// Increment is
type Increment struct {
	ID uint `json:"id"`
}

// UUID is
type UUID struct {
	ID string `json:"id"`
}

// Timestamps is
type Timestamps struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
