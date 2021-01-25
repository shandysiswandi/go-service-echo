package route

import (
	"go-rest-echo/config"
	"go-rest-echo/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

// HomeRoute is
func HomeRoute(e *echo.Echo, config *config.Config, db *db.Database) {
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Welcome to our API"})
	})
}
