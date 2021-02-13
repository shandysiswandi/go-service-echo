package sentrylib_test

import (
	"errors"
	"go-rest-echo/app/library/sentrylib"
	"go-rest-echo/config"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestSentry_CaptureMessage_And_CaptureError(t *testing.T) {
	is := assert.New(t)

	if err := godotenv.Load(".sentry"); err != nil {
		is.Nil(err)
	}

	s := sentrylib.New(config.New())
	msg := s.CaptureMessage("message")
	exc := s.CaptureError(errors.New("error"))

	is.NotNil(s)
	is.NotNil(msg)
	is.NotNil(exc)
}
