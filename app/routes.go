package app

import (
	"go-rest-echo/config"
	"go-rest-echo/db"
	"go-rest-echo/internal/tasks"
	"go-rest-echo/internal/users"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func routeWithoutJwt(e *echo.Echo, c *config.Config) {
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Welcome to our API"})
	})
}

func routeWithJwt(e *echo.Echo, c *config.Config, db *db.Database) {
	api := e.Group("/api")
	api.Use(middleware.JWT(c.JwtSecret))

	// define variables and inject
	var (
		r *echo.Group
		// tasks
		taskRepo     = tasks.NewMysql(db)
		taskUsecase  = tasks.NewUsecase(taskRepo)
		taskDelivery = tasks.NewDelivery(taskUsecase)

		// users
		userRepo     = users.NewMysql(db)
		userUsecase  = users.NewUsecase(userRepo)
		userDelivery = users.NewDelivery(userUsecase)
	)

	r = api.Group("/tasks")
	r.GET("", taskDelivery.Fetch)
	r.GET("/:id", taskDelivery.Get)
	r.POST("", taskDelivery.Create)
	r.PUT("/:id", taskDelivery.Update)
	r.DELETE("/:id", taskDelivery.Delete)

	r = api.Group("/users")
	r.GET("", userDelivery.Fetch)
	r.GET("/:id", userDelivery.Get)
	r.POST("", userDelivery.Create)
	r.PUT("/:id", userDelivery.Update)
	r.DELETE("/:id", userDelivery.Delete)
}