package route

import (
	d "go-rest-echo/delivery/http/task"
	r "go-rest-echo/repository/task"
	u "go-rest-echo/usecase/task"

	"github.com/labstack/echo/v4"
)

// TaskRoute is
func TaskRoute(e *echo.Echo) {
	// define variables and inject
	repository := r.NewMysql()
	usecase := u.NewUsecase(repository)
	delivery := d.NewDelivery(usecase)

	// create group and route
	r := e.Group("/tasks")
	r.GET("", delivery.Fetch)
	r.GET("/:id", delivery.Get)
	r.POST("", delivery.Create)
	r.PUT("/:id", delivery.Update)
	r.DELETE("/:id", delivery.Delete)
}
