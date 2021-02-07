package service_test

import (
	"go-rest-echo/config"
	"go-rest-echo/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	c := config.NewConfiguration()
	service := service.New(c)

	assert.NotNil(t, service)
	assert.NotNil(t, service.JWT)
	assert.NotNil(t, service.Redis)
	assert.Nil(t, service.Sentry)
}
