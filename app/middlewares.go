package app

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Middlewares is
func Middlewares(e *echo.Echo) {
	// Remove Trailing Slash
	e.Pre(middleware.RemoveTrailingSlash())

	// logger
	logger(e)

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

func logger(e *echo.Echo) {
	time := "\033[33mTIME\033[0m: ${time_rfc3339}"
	method := "\033[33mMETHOD\033[0m: ${method}"
	uri := "\033[33mURI\033[0m: ${uri}"
	ip := "\033[33mIP\033[0m: ${remote_ip}"
	status := "\033[33mSTATUS\033[0m: ${status}"
	latency := "\033[33mLATENCY\033[0m: ${latency_human}"

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: fmt.Sprintf("%s | %s | %s | %s | %s | %s\n", time, method, uri, ip, status, latency),
	}))
}
