package app

import (
	"net/http"

	taskHandler "go-rest-echo/domain/task/delivery"
	userHandler "go-rest-echo/domain/user/delivery"

	"github.com/labstack/echo/v4"
)

// Route is
func Route(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Auto Migren Pala Gue")
	})

	task := e.Group("/tasks")
	task.GET("", taskHandler.FetchTask)
	task.GET("/:id", taskHandler.GetTask)
	task.POST("", taskHandler.CreateTask)
	task.PUT("/:id", taskHandler.UpdateTask)
	task.DELETE("/:id", taskHandler.DeleteTask)

	user := e.Group("/users")
	user.GET("", userHandler.FetchUser)
	user.GET("/:id", userHandler.GetUser)
	user.POST("", userHandler.CreateUser)
	user.PUT("/:id", userHandler.UpdateUser)
	user.DELETE("/:id", userHandler.DeleteUser)
}
