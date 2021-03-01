package routes

import (
	"go-service-echo/app/library/redis"
	"go-service-echo/app/library/sentry"
	"go-service-echo/app/library/token"
	"go-service-echo/domain"
	"go-service-echo/domain/authentication"
	"go-service-echo/domain/users"
	"go-service-echo/infrastructure/database"
	"go-service-echo/infrastructure/jsonplaceholder"

	"github.com/labstack/echo/v4"
)

// Routes is
type Routes struct {
	engine *echo.Echo
	prefix *echo.Group
}

// New is
func New(e *echo.Echo) *Routes {
	p := e.Group("/api")
	return &Routes{e, p}
}

// Default is
func (r *Routes) Default(db *database.Database, tok *token.Token, red *redis.Redis, sen *sentry.Sentry, jph *jsonplaceholder.JSONPlaceHolder) *Routes {
	dHanlder := domain.NewDefaultHandler(db, tok, red, sen, jph)

	r.engine.GET("/", dHanlder.Default)                                  // default route and check | db | token | sentry
	r.engine.Any("/cors", dHanlder.CORS)                                 // for cors testing
	r.engine.GET("/favicon.ico", dHanlder.Favicon)                       // for request via browser
	r.engine.GET("/example-external-call", dHanlder.ExampleExternalCall) // example external call

	return r
}

// Auth is
func (r *Routes) Auth(db *database.Database, tok *token.Token) *Routes {
	userRepo := users.NewGormRepository(db)
	authUsecase := authentication.NewUsecase(userRepo, tok)
	authHanlder := authentication.NewHandler(authUsecase)

	authRoute := r.engine.Group("/auth")
	authRoute.POST("/login", authHanlder.Login)
	authRoute.POST("/register", authHanlder.Login)
	authRoute.POST("/forgot-password", authHanlder.Login)
	authRoute.POST("/reset-password", authHanlder.Login)
	authRoute.POST("/refresh", authHanlder.Login)
	authRoute.POST("/logout", authHanlder.Login)

	return r
}

// Users is
func (r *Routes) Users(db *database.Database) *Routes {
	userRepo := users.NewGormRepository(db)
	userUsecase := users.NewUserUsecase(userRepo)
	userHanlder := users.NewUserHandler(userUsecase)

	usersRoute := r.prefix.Group("/users")
	usersRoute.GET("", userHanlder.Fetch)
	usersRoute.GET("/:id", userHanlder.Get)
	usersRoute.POST("", userHanlder.Create)
	usersRoute.PUT("/:id", userHanlder.Update)
	usersRoute.DELETE("/:id", userHanlder.Delete)

	return r
}
