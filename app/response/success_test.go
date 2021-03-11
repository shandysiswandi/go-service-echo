package response_test

import (
	"go-service-echo/app/response"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func Test_NewSuccess(t *testing.T) {
	c := echo.New().NewContext(httptest.NewRequest(echo.GET, "/", nil), httptest.NewRecorder())
	act := response.NewSuccess(c, 200, "", nil)
	assert.Nil(t, act)
}
