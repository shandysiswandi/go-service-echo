package route

import (
	"go-rest-echo/config"
	"go-rest-echo/db"
	"go-rest-echo/internal/tasks"

	"github.com/labstack/echo/v4"
)

// TaskRoute is
func TaskRoute(e *echo.Echo, config *config.Config, db *db.Database) {
	// define variables and inject
	repository := tasks.NewMysql(db)
	usecase := tasks.NewUsecase(repository)
	delivery := tasks.NewDelivery(usecase)

	// create group and route
	r := e.Group("/tasks")
	r.GET("", delivery.Fetch)
	r.GET("/:id", delivery.Get)
	r.POST("", delivery.Create)
	r.PUT("/:id", delivery.Update)
	r.DELETE("/:id", delivery.Delete)
}
