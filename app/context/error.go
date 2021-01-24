package context

import (
	"errors"
	"go-rest-echo/util"
	"net/http"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type responseError struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}

// BadRequest is | 400
func (c *CustomContext) BadRequest(err interface{}) error {
	return c.JSON(http.StatusBadRequest, responseError{
		Status:  false,
		Message: "Bad Request, something wrong on your request",
		Error:   err,
	})
}

// NotFound is | 404
func (c *CustomContext) NotFound(err interface{}) error {
	return c.JSON(http.StatusNotFound, responseError{
		Status:  false,
		Message: "Not Found, your request data not found in our database",
		Error:   err,
	})
}

// UnprocessableEntity is | 422
func (c *CustomContext) UnprocessableEntity(err interface{}) error {
	var e []map[string]interface{}

	for _, err := range err.(validator.ValidationErrors) {
		e = append(e, map[string]interface{}{
			"key":     util.SnakeCase(err.StructField()),
			"message": getMessageValidation(err),
			"value":   err.Value(),
		})
	}

	return c.JSON(http.StatusUnprocessableEntity, responseError{
		Status:  false,
		Message: "Validation Failed",
		Error:   e,
	})
}

// private function to get message validation by flag | tag | reflect
func getMessageValidation(e validator.FieldError) (msg string) {
	switch e.Tag() {
	case "min":
		msg = "value must be greater than"
	case "required":
		msg = "value must be required"
	default:
		msg = "value must be validate"
	}

	if e.Param() != "" {
		msg = msg + " " + e.Param()
	}

	return msg
}

// InternalServerError is | 500
func (c *CustomContext) InternalServerError(err interface{}) error {
	return c.JSON(http.StatusInternalServerError, responseError{
		Status:  false,
		Message: "Internal Server Error",
		Error:   err,
	})
}

// HandleErrors is | 40x - 50x
func (c *CustomContext) HandleErrors(e error) error {
	if errors.Is(e, gorm.ErrRecordNotFound) {
		return c.NotFound(gorm.ErrRecordNotFound)
	}

	if errors.Is(e, gorm.ErrInvalidTransaction) {
		return c.InternalServerError(gorm.ErrInvalidTransaction)
	}

	if errors.Is(e, gorm.ErrNotImplemented) {
		return c.InternalServerError(gorm.ErrNotImplemented)
	}

	if errors.Is(e, gorm.ErrMissingWhereClause) {
		return c.InternalServerError(gorm.ErrMissingWhereClause)
	}

	if errors.Is(e, gorm.ErrUnsupportedRelation) {
		return c.InternalServerError(gorm.ErrUnsupportedRelation)
	}

	if errors.Is(e, gorm.ErrPrimaryKeyRequired) {
		return c.InternalServerError(gorm.ErrPrimaryKeyRequired)
	}

	if errors.Is(e, gorm.ErrModelValueRequired) {
		return c.InternalServerError(gorm.ErrModelValueRequired)
	}

	if errors.Is(e, gorm.ErrInvalidData) {
		return c.InternalServerError(gorm.ErrInvalidData)
	}

	if errors.Is(e, gorm.ErrUnsupportedDriver) {
		return c.InternalServerError(gorm.ErrUnsupportedDriver)
	}

	if errors.Is(e, gorm.ErrRegistered) {
		return c.InternalServerError(gorm.ErrRegistered)
	}

	if errors.Is(e, gorm.ErrInvalidField) {
		return c.InternalServerError(gorm.ErrInvalidField)
	}

	if errors.Is(e, gorm.ErrEmptySlice) {
		return c.InternalServerError(gorm.ErrEmptySlice)
	}

	if errors.Is(e, gorm.ErrDryRunModeUnsupported) {
		return c.InternalServerError(gorm.ErrDryRunModeUnsupported)
	}

	return c.InternalServerError(e)
}
