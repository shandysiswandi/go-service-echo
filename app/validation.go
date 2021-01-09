package app

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// Validation is
type valid struct {
	validator *validator.Validate
}

// Validate is
func (v *valid) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

// Validation is
func Validation(e *echo.Echo) {
	e.Validator = &valid{validator: validator.New()}
}
