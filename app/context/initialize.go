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
	e.HTTPErrorHandler = httpErrorHandler
}

func httpErrorHandler(e error, c echo.Context) {
	code := http.StatusInternalServerError
	message := "Internal Server Error"

	if he, ok := e.(*echo.HTTPError); ok {
		switch he.Code {
		case 404:
			message = "The URL you want is not in this application."
			break
		case 405:
			message = "The URL you want is not using this METHOD."
			break
		}

		code = he.Code
	}

	c.JSON(code, responseError{
		Success: false,
		Message: message,
		Error:   e,
	})
}
