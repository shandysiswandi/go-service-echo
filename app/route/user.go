package route

import (
	"go-rest-echo/db"
	d "go-rest-echo/delivery/http/user"
	r "go-rest-echo/repository/user"
	u "go-rest-echo/usecase/user"

	"github.com/labstack/echo/v4"
)

// UserRoute is
func UserRoute(e *echo.Echo, db *db.Database) {
	// define variables and inject
	repository := r.NewMysql(db)
	usecase := u.NewUsecase(repository)
	delivery := d.NewDelivery(usecase)

	// create group and route
	r := e.Group("/users")
	r.GET("", delivery.Fetch)
	r.GET("/:id", delivery.Get)
	r.POST("", delivery.Create)
	r.PUT("/:id", delivery.Update)
	r.DELETE("/:id", delivery.Delete)
}
