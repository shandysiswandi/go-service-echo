package app

import (
	"go-service-echo/app/context"
	"go-service-echo/app/library/redis"
	"go-service-echo/app/library/sentry"
	"go-service-echo/app/library/token"
	"go-service-echo/app/middlewares"
	"go-service-echo/app/routes"
	"go-service-echo/app/validation"
	"go-service-echo/config"
	"go-service-echo/infrastructure/database"
	"go-service-echo/infrastructure/jsonplaceholder"
	"go-service-echo/util/logger"

	"github.com/labstack/echo/v4"
)

// App is
type App struct {
	engine          *echo.Echo
	config          *config.Config
	database        *database.Database
	token           *token.Token
	sentry          *sentry.Sentry
	redis           *redis.Redis
	jsonplaceholder *jsonplaceholder.JSONPlaceHolder
}

// New is
func New(config *config.Config) *App {
	/********** ********** ********** ********** **********/
	/* create new echo engine
	/********** ********** ********** ********** **********/
	engine := echo.New()

	/********** ********** ********** ********** **********/
	/* create new database variable
	/********** ********** ********** ********** **********/
	database, err := database.New(config.Database)
	if err != nil {
		logger.Error(err)
	}

	/********** ********** ********** ********** **********/
	/* create new library token variable
	/********** ********** ********** ********** **********/
	token, err := token.New(config.Token)
	if err != nil {
		logger.Error(err)
	}

	/********** ********** ********** ********** **********/
	/* create new library sentry variable
	/********** ********** ********** ********** **********/
	sentry, err := sentry.New(config.Sentry)
	if err != nil {
		logger.Error(err)
	}

	/********** ********** ********** ********** **********/
	/* create new library redis variable
	/********** ********** ********** ********** **********/
	redis := redis.New(config.Redis)

	/********** ********** ********** ********** **********/
	/* create new external call jsonplaceholder variable
	/********** ********** ********** ********** **********/
	jsonPlaceHolder := jsonplaceholder.New(config.External.JSONPlaceHolder)

	/********** ********** ********** ********** **********/
	/* return
	/********** ********** ********** ********** **********/
	return &App{
		engine,
		config,
		database,
		token,
		sentry,
		redis,
		jsonPlaceHolder,
	}
}

// RegisterContext is
func (app *App) RegisterContext() *App {
	context.New(app.engine)
	return app
}

// RegisterValidation is
func (app *App) RegisterValidation() *App {
	validation.New(app.engine)
	return app
}

// RegisterMiddlewares is
func (app *App) RegisterMiddlewares() *App {
	middlewares.New(app.engine).
		PreRouter().
		PraRouter(app.token)

	return app
}

// RegisterRoutes is
func (app *App) RegisterRoutes() *App {
	routes.New(app.engine).
		Default(app.database, app.token, app.redis, app.sentry, app.jsonplaceholder).
		Auth(app.database, app.token).
		Users(app.database)

	return app
}

// GetEngine is
func (app *App) GetEngine() *echo.Echo {
	return app.engine
}

// GetLogger is
func (app *App) GetLogger() echo.Logger {
	return app.engine.Logger
}
