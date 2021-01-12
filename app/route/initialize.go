package route

import (
	"fmt"
	"net/http"

	"go-rest-echo/domain/task"
	"go-rest-echo/domain/user"
	"go-rest-echo/helper"

	"github.com/labstack/echo/v4"
)

// Initialize is
func Initialize(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("%s ----> %s", helper.Env("NAME"), helper.Env("ENV")))
	})

	task.NewRouter().Initialize(e, "/tasks")
	user.NewRouter().Initialize(e, "/users")
}
