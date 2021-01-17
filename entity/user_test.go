package entity_test

import (
	"go-rest-echo/entity"
	"testing"
)

func TestUser(t *testing.T) {
	t.Run("1 Table Name", func(t *testing.T) {
		user := new(entity.User)
		got := user.TableName()
		want := "users"

		if got != want {
			t.Errorf("Expected `%v`, but got %v", want, got)
		}
	})
}
