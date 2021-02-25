package jsonplaceholder_test

import (
	"go-service-echo/config"
	"go-service-echo/external/jsonplaceholder"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestNew_With_URL_Empty(t *testing.T) {
	assert.Nil(t, jsonplaceholder.New(""))
}

func TestNew_With_URL_From_ENV(t *testing.T) {
	is := assert.New(t)

	if err := godotenv.Load(".env.test"); err != nil {
		is.Nil(err)
	}

	actual := jsonplaceholder.New(config.New().External.JSONPlaceHolder)

	is.NotNil(actual)
}
