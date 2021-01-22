package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Gzip is
func Gzip(e *echo.Echo) {
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{Level: 5}))
}
