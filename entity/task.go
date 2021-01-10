package entity

import (
	"go-rest-echo/app/base"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TaskTable is
const TaskTable string = "tasks"

// Task is
type Task struct {
	base.UUID
	UserID      string `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	base.Timestamp
}

/*
 * Hooks GORM
 */

// TableName is
func (Task) TableName() string {
	return "tasks"
}

// BeforeCreate is
func (u *Task) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return nil
}
