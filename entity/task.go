package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Task is
type Task struct {
	UUID
	UserID      string `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	Timestamps
}

// TableName is
func (Task) TableName() string {
	return "tasks"
}

// BeforeCreate is
func (u *Task) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return nil
}
