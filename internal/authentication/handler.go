package authentication

import (
	"go-service-echo/app/context"
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
func (w *AuthHandler) Login(cc echo.Context) error {
	c := cc.(*context.CustomContext)

	// define variables
	pl := PayloadLogin{}

	// binding
	if err := c.Bind(&pl); err != nil {
		return c.BadRequest(err)
	}

	// validation
	if err := c.Validate(&pl); err != nil {
		return c.UnprocessableEntity(err)
	}

	// usecase
	result, err := w.usecase.Login(&pl)
	if err != nil {
		return c.HandleErrors(err)
	}

	// response
	return c.Success(http.StatusOK, "auth login", result)
}
