package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// RemoveTrailingSlash is
func RemoveTrailingSlash(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())
}
