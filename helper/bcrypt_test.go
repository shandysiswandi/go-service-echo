package helper_test

import "testing"

const MESSAGE string = "Expected `%s`, but got %s"

func TestHashPassword(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		got := ""
		want := ""

		if got != want {
			t.Errorf(MESSAGE, want, got)
		}
	})
}

func TestCheckPasswordHash(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		got := ""
		want := ""

		if got != want {
			t.Errorf(MESSAGE, want, got)
		}
	})
}
