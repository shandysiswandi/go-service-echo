package helper_test

import (
	"go-rest-echo/helper"
	"testing"
)

func TestSnakeCase(t *testing.T) {
	t.Run("1 All text is capitalize only one word", func(t *testing.T) {
		got := helper.SnakeCase("CAPITALIZE")
		want := "capitalize"

		if got != want {
			t.Errorf("Expected `%s`, but got %s", want, got)
		}
	})

	t.Run("2 All text are capitalize with two words", func(t *testing.T) {
		got := helper.SnakeCase("CAPITALIZE TEXT")
		want := "capitalize_text"

		if got != want {
			t.Errorf("Expected `%s`, but got %s", want, got)
		}
	})

	t.Run("3 All text is lower with one word", func(t *testing.T) {
		got := "lower"
		want := "lower"

		if got != want {
			t.Errorf("Expected `%s`, but got %s", want, got)
		}
	})

	t.Run("3 All text are lower with two words", func(t *testing.T) {
		got := "lower case"
		want := "lower_case"

		if got != want {
			t.Errorf("Expected `%s`, but got %s", want, got)
		}
	})
}
