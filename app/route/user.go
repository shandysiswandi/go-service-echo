package route

import (
	"go-rest-echo/config"
	"go-rest-echo/db"
	"go-rest-echo/internal/users"

	"github.com/labstack/echo/v4"
)

// UserRoute is
func UserRoute(e *echo.Echo, config *config.Config, db *db.Database) {
	// define variables and inject
	repository := users.NewMysql(db)
	usecase := users.NewUsecase(repository)
	delivery := users.NewDelivery(usecase)

	// create group and route
	r := e.Group("/users")
	r.GET("", delivery.Fetch)
	r.GET("/:id", delivery.Get)
	r.POST("", delivery.Create)
	r.PUT("/:id", delivery.Update)
	r.DELETE("/:id", delivery.Delete)
}
