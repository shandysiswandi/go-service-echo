package context

import (
	"errors"
	"go-rest-echo/helper"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type responseSuccess struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type responseError struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}

// CustomContext is
type CustomContext struct {
	echo.Context
}

// Initialize is
func Initialize(e *echo.Echo) {
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &CustomContext{c}
			return next(cc)
		}
	})
}

/* receiver & function base on custom context */

// Success is
func (c *CustomContext) Success(s int, m string, d interface{}) error {
	return c.JSON(s, responseSuccess{
		Status:  true,
		Message: m,
		Data:    d,
	})
}

// BadRequest is
func (c *CustomContext) BadRequest(err interface{}) error {
	return c.JSON(http.StatusBadRequest, responseError{
		Status:  false,
		Message: "Bad Request, something wrong on your request",
		Error:   err,
	})
}

// NotFound is
func (c *CustomContext) NotFound(err interface{}) error {
	return c.JSON(http.StatusNotFound, responseError{
		Status:  false,
		Message: "Not Found, your request data not found in our database",
		Error:   err,
	})
}

// ValidationFail is
func (c *CustomContext) ValidationFail(err interface{}) error {
	var e []map[string]interface{}

	for _, err := range err.(validator.ValidationErrors) {
		e = append(e, map[string]interface{}{
			"key":     helper.SnakeCase(err.StructField()),
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

// InternalServerError is
func (c *CustomContext) InternalServerError(err interface{}) error {
	return c.JSON(http.StatusInternalServerError, responseError{
		Status:  false,
		Message: "Internal Server Error",
		Error:   err,
	})
}

// HandleErrors is
func (c *CustomContext) HandleErrors(e error) error {
	if errors.Is(e, gorm.ErrRecordNotFound) {
		return c.NotFound(gorm.ErrRecordNotFound)
	}

	if errors.Is(e, gorm.ErrInvalidTransaction) {
		return c.NotFound(gorm.ErrInvalidTransaction)
	}

	if errors.Is(e, gorm.ErrNotImplemented) {
		return c.NotFound(gorm.ErrNotImplemented)
	}

	if errors.Is(e, gorm.ErrMissingWhereClause) {
		return c.NotFound(gorm.ErrMissingWhereClause)
	}

	if errors.Is(e, gorm.ErrUnsupportedRelation) {
		return c.NotFound(gorm.ErrUnsupportedRelation)
	}

	if errors.Is(e, gorm.ErrPrimaryKeyRequired) {
		return c.NotFound(gorm.ErrPrimaryKeyRequired)
	}

	if errors.Is(e, gorm.ErrModelValueRequired) {
		return c.NotFound(gorm.ErrModelValueRequired)
	}

	if errors.Is(e, gorm.ErrInvalidData) {
		return c.NotFound(gorm.ErrInvalidData)
	}

	if errors.Is(e, gorm.ErrUnsupportedDriver) {
		return c.NotFound(gorm.ErrUnsupportedDriver)
	}

	if errors.Is(e, gorm.ErrRegistered) {
		return c.NotFound(gorm.ErrRegistered)
	}

	if errors.Is(e, gorm.ErrInvalidField) {
		return c.NotFound(gorm.ErrInvalidField)
	}

	if errors.Is(e, gorm.ErrEmptySlice) {
		return c.NotFound(gorm.ErrEmptySlice)
	}

	if errors.Is(e, gorm.ErrDryRunModeUnsupported) {
		return c.NotFound(gorm.ErrDryRunModeUnsupported)
	}

	return c.InternalServerError(e)
}

/* private function */
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
