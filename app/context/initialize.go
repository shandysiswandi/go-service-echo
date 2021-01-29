package context

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// CustomContext is
type CustomContext struct {
	echo.Context
}

// NewCustomContext is
func NewCustomContext(e *echo.Echo) {
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error { return next(&CustomContext{c}) }
	})

	// set custom error
	e.HTTPErrorHandler = func(e error, c echo.Context) {
		code := http.StatusInternalServerError
		message := "Internal Server Error"

		if he, ok := e.(*echo.HTTPError); ok {
			code = he.Code
			message = he.Message.(string)
		}

		c.JSON(code, responseError{
			Status:  false,
			Message: message,
			Error:   e,
		})
	}
}
