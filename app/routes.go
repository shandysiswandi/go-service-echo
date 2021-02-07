package app

import (
	"go-rest-echo/config"
	"go-rest-echo/db"
	"go-rest-echo/external"
	"go-rest-echo/internal/authentication"
	"go-rest-echo/internal/blogs"
	"go-rest-echo/internal/tasks"
	"go-rest-echo/internal/users"
	"go-rest-echo/internal/welcomes"
	"go-rest-echo/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func routes(e *echo.Echo, c *config.Config, db *db.Database, s *service.Service, ex *external.External) {

	var (
		// welcomes
		welcomeUsecase  = welcomes.NewUsecase()
		welcomeDelivery = welcomes.NewWeb(welcomeUsecase)

		// users
		userRepo     = users.NewMysql(db)
		userUsecase  = users.NewUsecase(userRepo)
		userDelivery = users.NewDelivery(userUsecase)

		// auth
		authUsecase  = authentication.NewUsecase(userRepo, s)
		authDelivery = authentication.NewWeb(authUsecase)

		// tasks
		taskRepo     = tasks.NewMysql(db)
		taskUsecase  = tasks.NewUsecase(taskRepo)
		taskDelivery = tasks.NewDelivery(taskUsecase)

		// blogs
		blogRepo     = blogs.NewMysql(db)
		blogUsecase  = blogs.NewUsecase(blogRepo)
		blogDelivery = blogs.NewWeb(blogUsecase)
	)

	e.GET("/", welcomeDelivery.Home)
	e.GET("/monitor-database", welcomeDelivery.MonitorDatabase)
	e.GET("/monitor-service", welcomeDelivery.MonitorService)

	r := e.Group("/auth")
	r.POST("/login", authDelivery.Login)

	/******--Restricted--*****/
	api := e.Group("/api")
	api.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &service.JWTClaim{},
		SigningKey: c.Service.JWT.AccessSecret,
	}))

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

	r = api.Group("/blogs")
	r.GET("", blogDelivery.Fetch)
	r.GET("/:id", blogDelivery.Get)
	r.POST("", blogDelivery.Create)
	r.PUT("/:id", blogDelivery.Update)
	r.PATCH("/:id", blogDelivery.UpdateField)
	r.DELETE("/:id", blogDelivery.Delete)
}
