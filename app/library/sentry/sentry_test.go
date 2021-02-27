package sentry_test

import (
	"errors"
	"go-service-echo/app/library/sentry"
	"go-service-echo/config"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSentry_CaptureMessage_And_CaptureError(t *testing.T) {
	is := assert.New(t)

	os.Setenv("SENTRY_DSN", "")
	os.Setenv("SENTRY_ENV", "")

	s, err := sentry.New(config.New().Sentry)
	is.Nil(err)

	msg := s.Message("message")
	exc := s.Error(errors.New("error"))

	is.NotNil(s)
	is.NotNil(msg)
	is.NotNil(exc)
}

func TestNew_Error(t *testing.T) {
	is := assert.New(t)

	os.Setenv("SENTRY_DSN", "https://a@a/a")
	os.Setenv("SENTRY_ENV", "development")

	s, err := sentry.New(config.New().Sentry)
	is.Error(err)
	is.Nil(s)
}
