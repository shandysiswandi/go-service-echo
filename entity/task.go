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
	UserID      string `json:"user_id" gorm:"type:varchar(36)" validate:"required"`
	Title       string `json:"title" gorm:"type:varchar(100); not null" validate:"required,min=5"`
	Description string `json:"description" gorm:"type:varchar(100); not null" validate:"required,min=15"`
	Completed   bool   `json:"completed"`
	base.Timestamp
}

// BeforeCreate is
func (u *Task) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return
}
