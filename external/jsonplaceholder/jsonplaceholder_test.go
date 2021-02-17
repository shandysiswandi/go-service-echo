package jsonplaceholder_test

import (
	"go-rest-echo/config"
	"go-rest-echo/external/jsonplaceholder"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestNew_With_URL_Empty(t *testing.T) {
	assert.Nil(t, jsonplaceholder.New(""))
}

func TestNew_With_URL_From_ENV(t *testing.T) {
	if err := godotenv.Load(".env"); err != nil {
		t.Error(".env Not Found")
	}

	is := assert.New(t)
	conf := config.New()
	actual := jsonplaceholder.New(conf.External.JsonplaceholderURL)

	is.NotNil(actual)
}
