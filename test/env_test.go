package test

import (
	"testing"

	. "go-rest-echo/helper"

	"github.com/joho/godotenv"
)

func TestEnv(t *testing.T) {
	err := godotenv.Load(".env.test")
	if err != nil {
		t.Errorf("No %v", err)
		return
	}

	t.Run("1. ENV value", func(t *testing.T) {
		got := Env("ENV")
		expected := "testing"

		if got != expected {
			t.Errorf("Expected `%s`, but got %s", expected, got)
		}
	})

	t.Run("2. PORT value", func(t *testing.T) {
		got := Env("PORT")
		expected := "3000"

		if got != expected {
			t.Errorf("Expected `%s`, but got %s", expected, got)
		}
	})

	t.Run("3. Empty string value", func(t *testing.T) {
		got := Env("")
		expected := ""

		if got != expected {
			t.Errorf("Expected `%s`, but got %s", expected, got)
		}
	})

	t.Run("4. Empty key value", func(t *testing.T) {
		got := Env("EMPTY_KEY")
		expected := ""

		if got != expected {
			t.Errorf("Expected `%s`, but got %s", expected, got)
		}
	})
}
