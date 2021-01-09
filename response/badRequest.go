package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// BadRequest is
func BadRequest(c echo.Context, err interface{}) error {
	return c.JSON(http.StatusBadRequest, Error{
		Status:  false,
		Message: "Bad Request, something wrong on your request",
		Error:   err,
	})
}
