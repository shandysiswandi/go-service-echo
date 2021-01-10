package task

import (
	"go-rest-echo/domain/task/create"
	"go-rest-echo/domain/task/delete"
	"go-rest-echo/domain/task/fetch"
	"go-rest-echo/domain/task/get"
	"go-rest-echo/domain/task/update"

	"github.com/labstack/echo/v4"
)

// Route is
func Route(e *echo.Echo) {
	r := e.Group("/users")
	r.GET("", fetch.Delivery)
	r.GET("/:id", get.Delivery)
	r.POST("", create.Delivery)
	r.PUT("/:id", update.Delivery)
	r.DELETE("/:id", delete.Delivery)
}
