package app

import (
	"go-service-echo/app/library/redis"
	"go-service-echo/app/library/sentry"
	"go-service-echo/app/library/token"
	"go-service-echo/app/middlewares"
	"go-service-echo/app/response"
	"go-service-echo/app/routes"
	"go-service-echo/app/validation"
	"go-service-echo/config"
	"go-service-echo/infrastructure/gormdb"
	"go-service-echo/infrastructure/jsonplaceholder"

	"github.com/labstack/echo/v4"
)

// App is
type App struct {
	engine          *echo.Echo
	config          *config.Config
	database        *gormdb.Database
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
	engine.HTTPErrorHandler = response.DefaultEchoError

	/********** ********** ********** ********** **********/
	/* create new database variable
	/********** ********** ********** ********** **********/
	database, err := gormdb.New(config.Database)
	if err != nil {
		println("Err: ", err)
	}

	/********** ********** ********** ********** **********/
	/* create new library token variable
	/********** ********** ********** ********** **********/
	token, err := token.New(config.Token)
	if err != nil {
		println("Err: ", err)
	}

	/********** ********** ********** ********** **********/
	/* create new library sentry variable
	/********** ********** ********** ********** **********/
	sentry, err := sentry.New(config.Sentry)
	if err != nil {
		println("Err: ", err)
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
	/* register validation
	/********** ********** ********** ********** **********/
	validation.New(engine)

	/********** ********** ********** ********** **********/
	/* register middlewares
	/********** ********** ********** ********** **********/
	middlewares.New(engine, token)

	/********** ********** ********** ********** **********/
	/* register routes
	/********** ********** ********** ********** **********/
	routes.New(&routes.Routes{
		Engine:          engine,
		Database:        database,
		Token:           token,
		Redis:           redis,
		Sentry:          sentry,
		JSONPlaceHolder: jsonPlaceHolder,
	})

	/********** ********** ********** ********** **********/
	/* return
	/********** ********** ********** ********** **********/
	return &App{engine, config, database, token, sentry, redis, jsonPlaceHolder}
}

// GetEngine is
func (app *App) GetEngine() *echo.Echo {
	return app.engine
}
