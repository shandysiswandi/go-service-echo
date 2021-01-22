package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Cors is
func Cors(e *echo.Echo) {
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
}
