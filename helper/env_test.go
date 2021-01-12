package helper_test

import (
	"go-rest-echo/helper"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/joho/godotenv"
)

func envPath() (b string) {
	_, b, _, _ = runtime.Caller(0)
	return filepath.Join(filepath.Dir(b), "../.env.test")
}

func TestEnv(t *testing.T) {
	err := godotenv.Load(envPath())
	if err != nil {
		t.Errorf("No %v", err)
		return
	}

	t.Run("1. ENV value", func(t *testing.T) {
		got := helper.Env("ENV")
		expected := "testing"

		if got != expected {
			t.Errorf("Expected `%s`, but got %s", expected, got)
		}
	})

	t.Run("2. PORT value", func(t *testing.T) {
		got := helper.Env("PORT")
		expected := "3000"

		if got != expected {
			t.Errorf("Expected `%s`, but got %s", expected, got)
		}
	})

	t.Run("3. Empty string value", func(t *testing.T) {
		got := helper.Env("")
		expected := ""

		if got != expected {
			t.Errorf("Expected `%s`, but got %s", expected, got)
		}
	})

	t.Run("4. Empty key value", func(t *testing.T) {
		got := helper.Env("EMPTY_KEY")
		expected := ""

		if got != expected {
			t.Errorf("Expected `%s`, but got %s", expected, got)
		}
	})
}
