package response

import (
	"go-rest-echo/helper"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// ValidationFail is
func ValidationFail(c echo.Context, err interface{}) error {
	var e []map[string]interface{}

	for _, err := range err.(validator.ValidationErrors) {
		e = append(e, map[string]interface{}{
			"key":     helper.SnakeCase(err.StructField()),
			"message": getMessage(err),
			"value":   err.Value(),
		})
	}

	return c.JSON(http.StatusUnprocessableEntity, Error{
		Status:  false,
		Message: "Validation Failed",
		Error:   e,
	})
}

func getMessage(e validator.FieldError) (msg string) {
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
