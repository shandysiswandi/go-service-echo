package service_test

import (
	"errors"
	"go-rest-echo/config"
	"go-rest-echo/service"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func init() {
	err := godotenv.Load(".env.test")
	if err != nil {
		return
	}
}

func TestSentry_CaptureMessage_And_CaptureError(t *testing.T) {
	c := config.New()
	sentry := service.NewSentry(c)
	msg := sentry.CaptureMessage("message")
	exc := sentry.CaptureError(errors.New("error"))

	assert.NotEqual(t, nil, sentry)
	assert.NotEqual(t, nil, msg)
	assert.NotEqual(t, nil, exc)
}
