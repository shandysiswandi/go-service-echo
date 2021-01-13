package helper_test

import (
	"go-rest-echo/helper"
	"testing"
)

const MESSAGE string = "Expected `%s`, but got %s"

func TestHashPassword(t *testing.T) {
	t.Run("1 Check five character from front", func(t *testing.T) {
		want := "$2a$10"
		got, err := helper.HashPassword("password")
		if err != nil {
			t.Errorf(MESSAGE, want, err.Error())
		}

		if got[0:6] != want {
			t.Errorf(MESSAGE, want, got)
		}
	})
}

func TestCheckPasswordHash(t *testing.T) {
	t.Run("1 Check password is valid", func(t *testing.T) {
		val, err := helper.HashPassword("password")
		if err != nil {
			t.Errorf(MESSAGE, "", err.Error())

		}
		want := true
		got := helper.CheckPasswordHash("password", val)

		if got != want {
			t.Errorf("Expected `%v`, but got %v", want, got)
		}
	})
}
