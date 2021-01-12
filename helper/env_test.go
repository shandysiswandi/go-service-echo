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
		want := "testing"

		if got != want {
			t.Errorf("Expected`%s`, but got %s", want, got)
		}
	})

	t.Run("2. PORT value", func(t *testing.T) {
		got := helper.Env("PORT")
		want := "3000"

		if got != want {
			t.Errorf("Expected`%s`, but got %s", want, got)
		}
	})

	t.Run("3. Empty string value", func(t *testing.T) {
		got := helper.Env("")
		want := ""

		if got != want {
			t.Errorf("Expected`%s`, but got %s", want, got)
		}
	})

	t.Run("4. Empty key value", func(t *testing.T) {
		got := helper.Env("EMPTY_KEY")
		want := ""

		if got != want {
			t.Errorf("Expected`%s`, but got %s", want, got)
		}
	})
}
