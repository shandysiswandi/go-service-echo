package route

import (
	"go-rest-echo/config"
	"go-rest-echo/db"
	d "go-rest-echo/delivery/task"
	r "go-rest-echo/repository/task"
	u "go-rest-echo/usecase/task"

	"github.com/labstack/echo/v4"
)

// TaskRoute is
func TaskRoute(e *echo.Echo, config *config.Config, db *db.Database) {
	// define variables and inject
	repository := r.NewMysql(db)
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
