package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Middlewares is
func Middlewares(e *echo.Echo) {
	// Remove Trailing Slash
	e.Pre(middleware.RemoveTrailingSlash())

	// logger
	e.Use(middleware.Logger())

	// recover panic
	e.Use(middleware.Recover())

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// body limit
	e.Use(middleware.BodyLimit("8M"))

	// Gzip compress
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{Level: 5}))

	// secure
	e.Use(middleware.Secure())
}
