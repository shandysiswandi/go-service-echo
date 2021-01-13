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

	t.Run("2 All text is lower with one word", func(t *testing.T) {
		got := helper.SnakeCase("lower")
		want := "lower"

		if got != want {
			t.Errorf("Expected `%s`, but got %s", want, got)
		}
	})

	t.Run("3 Text input is IsUUID", func(t *testing.T) {
		got := helper.SnakeCase("IsUUID")
		want := "is_uuid"

		if got != want {
			t.Errorf("Expected `%s`, but got %s", want, got)
		}
	})

	t.Run("4 Text input is UserId", func(t *testing.T) {
		got := helper.SnakeCase("UserId")
		want := "user_id"

		if got != want {
			t.Errorf("Expected `%s`, but got %s", want, got)
		}
	})

	t.Run("5 Text input is theURL", func(t *testing.T) {
		got := helper.SnakeCase("theURL")
		want := "the_url"

		if got != want {
			t.Errorf("Expected `%s`, but got %s", want, got)
		}
	})
}
