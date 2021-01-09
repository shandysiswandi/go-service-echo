package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// NotFound is
func NotFound(c echo.Context, err interface{}) error {
	return c.JSON(http.StatusNotFound, Error{
		Status:  false,
		Message: "Not Found, your request data not found in our database",
		Error:   err,
	})
}
