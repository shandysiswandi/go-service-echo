package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

// NewRouter is
func NewRouter(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("%s ----> %s", os.Getenv("NAME"), os.Getenv("ENV")))
	})
}
