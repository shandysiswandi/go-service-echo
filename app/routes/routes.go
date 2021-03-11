package routes

import (
	"go-service-echo/app/library/redis"
	"go-service-echo/app/library/sentry"
	"go-service-echo/app/library/token"
	"go-service-echo/domain"
	"go-service-echo/domain/authentication"
	"go-service-echo/domain/users"
	"go-service-echo/infrastructure/gormdb"
	"go-service-echo/infrastructure/jsonplaceholder"

	"github.com/labstack/echo/v4"
)

type (
	// Routes is
	Routes struct {
		Engine          *echo.Echo
		Database        *gormdb.Database
		Token           *token.Token
		Redis           *redis.Redis
		Sentry          *sentry.Sentry
		JSONPlaceHolder *jsonplaceholder.JSONPlaceHolder
	}
)

// New is
func New(r *Routes) *Routes {
	/* define variables */
	api := r.Engine.Group("/api")
	db := r.Database
	token := r.Token
	redis := r.Redis
	sentry := r.Sentry
	jsonPlaceHolder := r.JSONPlaceHolder

	/* Route Default */
	dhc := &domain.DefaultHandlerConfig{Database: db, Token: token, Redis: redis, Sentry: sentry, JSONPlaceHolder: jsonPlaceHolder}
	dHanlder := domain.NewDefaultHandler(dhc)

	r.Engine.GET("/", dHanlder.Default)                                  // default route and check | db | token | sentry
	r.Engine.Any("/cors", dHanlder.CORS)                                 // for cors testing
	r.Engine.GET("/favicon.ico", dHanlder.Favicon)                       // for request via browser
	r.Engine.GET("/example-external-call", dHanlder.ExampleExternalCall) // example external call

	/* Route Auth */
	userRepo := users.NewGormRepository(db.SQL)
	authUsecase := authentication.NewUsecase(userRepo, token)
	authHanlder := authentication.NewHandler(authUsecase)

	authRoute := r.Engine.Group("/auth")
	authRoute.POST("/login", authHanlder.Login)
	authRoute.POST("/register", authHanlder.Login)
	authRoute.POST("/forgot-password", authHanlder.Login)
	authRoute.POST("/reset-password", authHanlder.Login)
	authRoute.POST("/refresh", authHanlder.Login)
	authRoute.POST("/logout", authHanlder.Login)

	/* Route Users */
	userUsecase := users.NewUserUsecase(userRepo)
	userHanlder := users.NewUserHandler(userUsecase)

	usersRoute := api.Group("/users")
	usersRoute.GET("", userHanlder.Fetch)
	usersRoute.GET("/:id", userHanlder.Get)
	usersRoute.POST("", userHanlder.Create)
	usersRoute.PUT("/:id", userHanlder.Update)
	usersRoute.DELETE("/:id", userHanlder.Delete)

	/* return */
	return nil
}
