package entity

import (
	"go-rest-echo/entity/base"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TaskTable is
const TaskTable string = "tasks"

// Task is
type Task struct {
	base.UUID
	UserID      string `gorm:"type:varchar(36)" json:"user_id" validate:"required"`
	Title       string `gorm:"type:varchar(100)" json:"title" validate:"required, min=5"`
	Description string `gorm:"type:varchar(100)" json:"description" validate:"required, min=15"`
	Completed   bool   `json:"completed"`
	base.Timestamp
}

// BeforeCreate is
func (u *Task) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return
}
