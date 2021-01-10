package user

import (
	"go-rest-echo/domain/user/create"
	"go-rest-echo/domain/user/delete"
	"go-rest-echo/domain/user/fetch"
	"go-rest-echo/domain/user/get"
	"go-rest-echo/domain/user/update"

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
