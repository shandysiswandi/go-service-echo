package middleware

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Initialize is
func Initialize(e *echo.Echo) {
	// Remove Trailing Slash
	e.Pre(middleware.RemoveTrailingSlash())

	// logger
	time := "\033[33mTIME\033[0m: ${time_rfc3339}"
	method := "\033[33mMETHOD\033[0m: ${method}"
	uri := "\033[33mURI\033[0m: ${uri}"
	ip := "\033[33mIP\033[0m: ${remote_ip}"
	status := "\033[33mSTATUS\033[0m: ${status}"
	err := "\033[33mERROR\033[0m: ${error}"
	latency := "\033[33mLATENCY\033[0m: ${latency_human}"

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: fmt.Sprintf("%s | %s | %s | %s | %s | %s | %s\n", time, method, uri, ip, status, err, latency),
	}))

	// recover panic
	e.Use(middleware.Recover())

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{
			echo.HeaderAccept,
			echo.HeaderAcceptEncoding,
			echo.HeaderAccessControlAllowHeaders,
			echo.HeaderAuthorization,
			echo.HeaderContentType,
			echo.HeaderContentLength,
			echo.HeaderOrigin,
			echo.HeaderXCSRFToken,
		},
		MaxAge: 86400,
	}))

	// body limit
	e.Use(middleware.BodyLimit("8M"))

	// Gzip compress
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{Level: 5}))

	// secure
	e.Use(middleware.Secure())
}
