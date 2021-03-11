package users_test

import (
	"go-service-echo/domain/users"
	"go-service-echo/mocks"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func Test_UserHandler_Fetch(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/users", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	rec := httptest.NewRecorder()
	eContext := e.NewContext(req, rec)
	// a := eContext.(*context.CustomContext)

	// log.Println(a)

	mockUserUsecase := new(mocks.UserUsecase)
	h := users.NewUserHandler(mockUserUsecase)
	log.Println(h)

	// Assertions
	if assert.NoError(t, h.Fetch(eContext)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, nil, rec.Body.String())
	}
}

func Test_UserHandler_Get(t *testing.T) {
	//
}

func Test_UserHandler_GetByEmail(t *testing.T) {
	//
}

func Test_UserHandler_Create(t *testing.T) {
	//
}

func Test_UserHandler_Update(t *testing.T) {
	//
}

func Test_UserHandler_Delete(t *testing.T) {
	//
}
