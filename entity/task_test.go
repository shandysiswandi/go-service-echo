package entity_test

import (
	"go-rest-echo/entity"
	"testing"

	"gorm.io/gorm"
)

func TestTask(t *testing.T) {
	t.Run("1 Table Name", func(t *testing.T) {
		task := new(entity.Task)
		got := task.TableName()
		want := "tasks"

		if got != want {
			t.Errorf("Expected `%v`, but got %v", want, got)
		}
	})

	t.Run("2 Before Create Hook", func(t *testing.T) {
		task := new(entity.Task)
		db := new(gorm.DB)
		got := task.BeforeCreate(db)

		if got != nil {
			t.Errorf("Expected `%v`, but got %v", nil, got)
		}
	})

	t.Run("3 Task Field", func(t *testing.T) {
		task := new(entity.Task)

		if task.ID != "" {
			t.Errorf("Expected `%v`, but got %v", "", task.ID)
		}

		if task.UserID != "" {
			t.Errorf("Expected `%v`, but got %v", "", task.UserID)
		}

		if task.Title != "" {
			t.Errorf("Expected `%v`, but got %v", "", task.Title)
		}

		if task.Description != "" {
			t.Errorf("Expected `%v`, but got %v", "", task.Description)
		}

		if task.Completed != false {
			t.Errorf("Expected `%v`, but got %v", "", task.Completed)
		}
	})
}
