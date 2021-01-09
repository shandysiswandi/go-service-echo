package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// InternalServerError is
func InternalServerError(c echo.Context, err interface{}) error {
	return c.JSON(http.StatusBadRequest, Error{
		Status:  false,
		Message: "Internal Server Error",
		Error:   err,
	})
}
