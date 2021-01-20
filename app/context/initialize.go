package context

import "github.com/labstack/echo/v4"

// CustomContext is
type CustomContext struct {
	echo.Context
}

// NewCustomContext is
func NewCustomContext(e *echo.Echo) {
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error { return next(&CustomContext{c}) }
	})
}
