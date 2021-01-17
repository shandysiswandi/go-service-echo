package entity_test

import (
	"go-rest-echo/entity"
	"testing"
)

func TestAnimal(t *testing.T) {
	t.Run("1 Table Name", func(t *testing.T) {
		animal := new(entity.Animal)
		got := animal.TableName()
		want := "animals"

		if want != got {
			t.Errorf("Expected `%v`, but got `%v`", want, got)
		}
	})
}
