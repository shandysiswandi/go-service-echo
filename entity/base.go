package entity

import (
	"time"

	"gorm.io/gorm"
)

// Increment is
type Increment struct {
	ID uint `gorm:"primaryKey" json:"id"`
}

// UUID is
type UUID struct {
	ID string `gorm:"primaryKey; type:varchar(36)" json:"id"`
}

// Timestamps is
type Timestamps struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
