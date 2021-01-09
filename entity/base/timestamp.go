package base

import (
	"time"

	"gorm.io/gorm"
)

// Timestamp is
type Timestamp struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
