package response

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// ValidationFail is
func ValidationFail(c echo.Context, err interface{}) error {
	var e []map[string]interface{}

	for _, err := range err.(validator.ValidationErrors) {
		e = append(e, map[string]interface{}{
			"key":     err.StructField(),
			"message": "must be " + err.Tag(),
			"value":   err.Value(),
		})
	}

	return c.JSON(http.StatusBadRequest, Error{
		Status:  false,
		Message: "Validation Failed",
		Error:   e,
	})
}
