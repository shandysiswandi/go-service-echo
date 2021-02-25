package app

import (
	"go-service-echo/app/library/jwt"
	"go-service-echo/app/library/redis"
	"go-service-echo/app/library/sentry"
	"go-service-echo/config"
	"go-service-echo/db"
	"go-service-echo/external/jsonplaceholder"
	"go-service-echo/internal/authentication"
	"go-service-echo/internal/blogs"
	"go-service-echo/internal/tasks"
	"go-service-echo/internal/users"
	"go-service-echo/internal/welcomes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func routes(e *echo.Echo, c *config.Config, db *db.Database) {
	var (
		// library
		jwtLib = jwt.New(c.JWT)
		redis  = redis.New(c.Redis)
		sentry = sentry.New(c)

		// external (thrid-party)
		jph = jsonplaceholder.New(c.External.JSONPlaceHolder)

		// welcomes
		welcomeDelivery = welcomes.NewWeb(db, jwtLib, redis, sentry, jph)

		// users
		userRepo     = users.NewMysql(db)
		userUsecase  = users.NewUsecase(userRepo)
		userDelivery = users.NewDelivery(userUsecase)

		// auth
		authUsecase  = authentication.NewUsecase(userRepo, jwtLib)
		authDelivery = authentication.NewWeb(authUsecase)

		// tasks
		taskRepo     = tasks.NewMysql(db)
		taskUsecase  = tasks.NewUsecase(taskRepo)
		taskDelivery = tasks.NewDelivery(taskUsecase)

		// blogs
		blogMysqlRepository = blogs.NewMysql(db)
		// blogPostgresqlRepository = blogs.NewPostgresql(db)
		blogUsecase  = blogs.NewUsecase(blogMysqlRepository)
		blogDelivery = blogs.NewWeb(blogUsecase)
	)

	e.GET("/", welcomeDelivery.Home)
	e.GET("/favicon.ico", welcomeDelivery.Favicon)
	e.GET("/check-database", welcomeDelivery.CheckDatabase)
	e.GET("/check-library", welcomeDelivery.CheckLibrary)
	e.GET("/check-external", welcomeDelivery.CheckExternal)

	r := e.Group("/auth")
	r.POST("/login", authDelivery.Login)
	r.POST("/register", authDelivery.Login)
	r.POST("/forgot-password", authDelivery.Login)
	r.POST("/reset-password", authDelivery.Login)
	r.POST("/refresh", authDelivery.Login)
	r.POST("/logout", authDelivery.Login)

	/******--Restricted--*****/
	api := e.Group("/api")
	api.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &jwt.Claim{},
		SigningKey: c.JWT.AccessSecret,
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
