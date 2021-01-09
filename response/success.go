package response

import (
	"github.com/labstack/echo/v4"
)

// Success is
func Success(c echo.Context, s int, m string, d interface{}) error {
	return c.JSON(s, Response{
		Status:  true,
		Message: m,
		Data:    d,
	})
}
