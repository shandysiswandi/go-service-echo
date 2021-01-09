package helper

import (
	"errors"
	"go-rest-echo/response"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// HandleErrors is
func HandleErrors(c echo.Context, e error) error {
	if errors.Is(e, gorm.ErrRecordNotFound) {
		return response.NotFound(c, gorm.ErrRecordNotFound)
	}

	if errors.Is(e, gorm.ErrInvalidTransaction) {
		return response.NotFound(c, gorm.ErrInvalidTransaction)
	}

	if errors.Is(e, gorm.ErrNotImplemented) {
		return response.NotFound(c, gorm.ErrNotImplemented)
	}

	return response.InternalServerError(c, e)
}
