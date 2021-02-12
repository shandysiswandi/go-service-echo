package external_test

import (
	"go-rest-echo/config"
	"go-rest-echo/external"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestNew_With_Config_Nil(t *testing.T) {
	assert.Nil(t, external.New(nil))
}

func TestNew_With_Config_From_ENV(t *testing.T) {
	if err := godotenv.Load(".env"); err != nil {
		t.Error(".env Not Found")
	}

	is := assert.New(t)
	actual := external.New(config.NewConfiguration())

	is.NotNil(actual)
	is.NotNil(actual.JSONPlaceHolder)
}
