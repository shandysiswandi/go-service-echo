package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type valid struct {
	validator *validator.Validate
}

// New is
func New(e *echo.Echo) {
	e.Validator = &valid{validator: validator.New()}
}

func (v *valid) Validate(i interface{}) error {
	return v.validator.Struct(i)
}
