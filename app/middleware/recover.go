package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Recover is
func Recover(e *echo.Echo) {
	e.Pre(middleware.Recover())
}
