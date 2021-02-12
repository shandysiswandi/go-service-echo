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
	if err := godotenv.Load(".env.test"); err != nil {
		return
	}

	s := sentrylib.New(config.New())
	msg := s.CaptureMessage("message")
	exc := s.CaptureError(errors.New("error"))

	assert.NotEqual(t, nil, s)
	assert.NotEqual(t, nil, msg)
	assert.NotEqual(t, nil, exc)
}
