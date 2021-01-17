package entity_test

import (
	"go-rest-echo/entity"
	"testing"
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
}
