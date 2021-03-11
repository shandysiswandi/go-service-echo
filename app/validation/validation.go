package validation

import (
	"fmt"
	"go-service-echo/app/response"

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

// ValidateVar is
func ValidateVar(value interface{}, tag string) map[string]interface{} {
	var v = validator.New()
	var e map[string]interface{}

	if err := v.Var(value, tag); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			e = map[string]interface{}{"message": fmt.Sprintf("`%v` %s", err.Value(), response.GetMessageValidation(err))}
		}

		return e
	}

	return nil
}

// func (c *CustomContext) GetTokenData() token.PayloadData {
// 	tokPayload := c.Get("user").(*token.Payload)
// 	return tokPayload.Data
// }
