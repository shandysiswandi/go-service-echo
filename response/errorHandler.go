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

	if errors.Is(e, gorm.ErrMissingWhereClause) {
		return NotFound(c, gorm.ErrMissingWhereClause)
	}

	if errors.Is(e, gorm.ErrUnsupportedRelation) {
		return NotFound(c, gorm.ErrUnsupportedRelation)
	}

	if errors.Is(e, gorm.ErrPrimaryKeyRequired) {
		return NotFound(c, gorm.ErrPrimaryKeyRequired)
	}

	if errors.Is(e, gorm.ErrModelValueRequired) {
		return NotFound(c, gorm.ErrModelValueRequired)
	}

	if errors.Is(e, gorm.ErrInvalidData) {
		return NotFound(c, gorm.ErrInvalidData)
	}

	if errors.Is(e, gorm.ErrUnsupportedDriver) {
		return NotFound(c, gorm.ErrUnsupportedDriver)
	}

	if errors.Is(e, gorm.ErrRegistered) {
		return NotFound(c, gorm.ErrRegistered)
	}

	if errors.Is(e, gorm.ErrInvalidField) {
		return NotFound(c, gorm.ErrInvalidField)
	}

	if errors.Is(e, gorm.ErrEmptySlice) {
		return NotFound(c, gorm.ErrEmptySlice)
	}

	if errors.Is(e, gorm.ErrDryRunModeUnsupported) {
		return NotFound(c, gorm.ErrDryRunModeUnsupported)
	}

	return InternalServerError(c, e)
}
