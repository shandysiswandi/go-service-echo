package response_test

import (
	"go-service-echo/app/library/token"
	"go-service-echo/app/response"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func Test_BadRequest(t *testing.T) {
	c := echo.New().NewContext(httptest.NewRequest(echo.GET, "/", nil), httptest.NewRecorder())
	act := response.BadRequest(c, nil)
	assert.Nil(t, act)
}

func Test_Unauthorized(t *testing.T) {
	c := echo.New().NewContext(httptest.NewRequest(echo.GET, "/", nil), httptest.NewRecorder())
	act := response.Unauthorized(c, nil)
	assert.Nil(t, act)
}

func Test_UnprocessableEntityVar(t *testing.T) {
	c := echo.New().NewContext(httptest.NewRequest(echo.GET, "/", nil), httptest.NewRecorder())
	act := response.UnprocessableEntityVar(c, nil)
	assert.Nil(t, act)
}

func Test_UnprocessableEntity(t *testing.T) {
	c := echo.New().NewContext(httptest.NewRequest(echo.GET, "/", nil), httptest.NewRecorder())

	var user struct {
		FirstName string `validate:"min=10"`
		LastName  string `validate:"required"`
		Age       uint8  `validate:"gte=5,lte=130"`
		Email     string `validate:"email"`
	}
	validate := validator.New()
	err := validate.Struct(user)

	act := response.UnprocessableEntity(c, err)
	assert.Nil(t, act)
}

func Test_HandleErrors(t *testing.T) {
	c := echo.New().NewContext(httptest.NewRequest(echo.GET, "/", nil), httptest.NewRecorder())

	var err error

	err = response.HandleErrors(c, token.ErrExpiredToken)
	assert.Nil(t, err)

	err = response.HandleErrors(c, token.ErrInvalidToken)
	assert.Nil(t, err)

	err = response.HandleErrors(c, gorm.ErrRecordNotFound)
	assert.Nil(t, err)

	err = response.HandleErrors(c, response.ErrInvalidCredential)
	assert.Nil(t, err)

	err = response.HandleErrors(c, response.ErrFailedGenerateToken)
	assert.Nil(t, err)

	err = response.HandleErrors(c, nil)
	assert.Nil(t, err)
}
