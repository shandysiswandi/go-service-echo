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

	s := sentry.New(config.New())
	msg := s.CaptureMessage("message")
	exc := s.CaptureError(errors.New("error"))

	is.NotNil(s)
	is.NotNil(msg)
	is.NotNil(exc)
}
