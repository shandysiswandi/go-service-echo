package jsonplaceholder_test

import (
	"go-service-echo/config"
	"go-service-echo/infrastructure/jsonplaceholder"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew_With_URL_Empty(t *testing.T) {
	assert.Nil(t, jsonplaceholder.New(""))
}

func TestNew_With_URL_From_ENV(t *testing.T) {
	is := assert.New(t)

	os.Setenv("URL_JSONPLACEHOLDER", "https://jsonplaceholder.com")

	actual := jsonplaceholder.New(config.New().External.JSONPlaceHolder)

	is.NotNil(actual)
}
