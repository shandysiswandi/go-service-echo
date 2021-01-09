package response

import (
	"errors"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// HandleErrors is
func HandleErrors(c echo.Context, e error) error {
	if errors.Is(e, gorm.ErrRecordNotFound) {
		return NotFound(c, gorm.ErrRecordNotFound)
	}

	if errors.Is(e, gorm.ErrInvalidTransaction) {
		return NotFound(c, gorm.ErrInvalidTransaction)
	}

	if errors.Is(e, gorm.ErrNotImplemented) {
		return NotFound(c, gorm.ErrNotImplemented)
	}

	return InternalServerError(c, e)
}
