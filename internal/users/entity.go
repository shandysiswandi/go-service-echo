package users

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User is
type User struct {
	ID        string         `json:"id"`
	Name      string         `json:"name" validate:"required,min=5"`
	Email     string         `json:"email" validate:"required,email,min=5"`
	Password  string         `json:"password" validate:"required,min=6"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	// Task      []Task
}

// TableName is
func (User) TableName() string {
	return "users"
}

// SetID is
func (u *User) SetID() {
	u.ID = uuid.New().String()
}
