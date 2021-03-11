package authentication

import (
	"go-service-echo/app/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

// AuthHandler is
type AuthHandler struct {
	usecase *AuthUsecase
}

// NewHandler is
func NewHandler(u *AuthUsecase) *AuthHandler {
	return &AuthHandler{u}
}

// Login is
func (w *AuthHandler) Login(c echo.Context) error {
	// define variables
	pl := PayloadLogin{}

	// binding
	if err := c.Bind(&pl); err != nil {
		return response.BadRequest(c, err)
	}

	// validation
	if err := c.Validate(&pl); err != nil {
		return response.UnprocessableEntity(c, err)
	}

	// usecase
	result, err := w.usecase.Login(&pl)
	if err != nil {
		return response.HandleErrors(c, err)
	}

	// response
	return response.NewSuccess(c, http.StatusOK, "auth login", result)
}
