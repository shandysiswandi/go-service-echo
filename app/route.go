package app

import (
	"net/http"

	tac "go-rest-echo/domain/task/create"
	tad "go-rest-echo/domain/task/delete"
	taf "go-rest-echo/domain/task/fetch"
	tag "go-rest-echo/domain/task/get"
	tau "go-rest-echo/domain/task/update"

	usc "go-rest-echo/domain/user/create"
	usd "go-rest-echo/domain/user/delete"
	usf "go-rest-echo/domain/user/fetch"
	usg "go-rest-echo/domain/user/get"
	usu "go-rest-echo/domain/user/update"

	"go-rest-echo/helper"

	"github.com/labstack/echo/v4"
)

// Route is
func Route(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "ENV "+helper.Env("ENV"))
	})

	task := e.Group("/tasks")
	task.GET("", taf.Delivery)
	task.GET("/:id", tag.Delivery)
	task.POST("", tac.Delivery)
	task.PUT("/:id", tau.Delivery)
	task.DELETE("/:id", tad.Delivery)

	user := e.Group("/users")
	user.GET("", usf.Delivery)
	user.GET("/:id", usg.Delivery)
	user.POST("", usc.Delivery)
	user.PUT("/:id", usu.Delivery)
	user.DELETE("/:id", usd.Delivery)
}
