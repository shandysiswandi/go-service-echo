package app

import (
	"go-service-echo/app/library/redis"
	"go-service-echo/app/library/sentry"
	"go-service-echo/app/library/token"
	"go-service-echo/config"
	"go-service-echo/db"
	"go-service-echo/external/jsonplaceholder"
	"go-service-echo/internal"
	"go-service-echo/internal/authentication"
	"go-service-echo/internal/users"
	"go-service-echo/util/logger"

	"github.com/labstack/echo/v4"
)

func routes(e *echo.Echo, c *config.Config, db *db.Database) {
	// library Token
	token, err := token.New(c.Token)
	if err != nil {
		logger.Error(err)
	}

	// library sentry
	sentry, err := sentry.New(c.Sentry)
	if err != nil {
		logger.Error(err)
	}

	var (
		redis = redis.New(c.Redis)

		// external (thrid-party)
		jph = jsonplaceholder.New(c.External.JSONPlaceHolder)

		// welcomes
		dDefault = internal.NewDefault(db, token, redis, sentry, jph)

		// users
		userRepo     = users.NewMysql(db)
		userUsecase  = users.NewUsecase(userRepo)
		userDelivery = users.NewDelivery(userUsecase)

		// auth
		authUsecase  = authentication.NewUsecase(userRepo, nil)
		authDelivery = authentication.NewWeb(authUsecase)
	)

	e.GET("/", dDefault.Default)                                  // default route and check | db | token | sentry
	e.Any("/cors", dDefault.CORS)                                 // for cors testing
	e.GET("/favicon.ico", dDefault.Favicon)                       // for request via browser
	e.GET("/example-external-call", dDefault.ExampleExternalCall) // example external call

	r := e.Group("/auth")
	r.POST("/login", authDelivery.Login)
	r.POST("/register", authDelivery.Login)
	r.POST("/forgot-password", authDelivery.Login)
	r.POST("/reset-password", authDelivery.Login)
	r.POST("/refresh", authDelivery.Login)
	r.POST("/logout", authDelivery.Login)

	r = e.Group("/users")
	r.GET("", userDelivery.Fetch)
	r.GET("/:id", userDelivery.Get)
	r.POST("", userDelivery.Create)
	r.PUT("/:id", userDelivery.Update)
	r.DELETE("/:id", userDelivery.Delete)
}
