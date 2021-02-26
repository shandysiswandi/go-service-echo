package app

import (
	"go-service-echo/app/context"
	"go-service-echo/app/validation"
	"go-service-echo/config"
	"go-service-echo/db"

	"github.com/labstack/echo/v4"
)

// App is
type App struct {
	engine   *echo.Echo
	config   *config.Config
	database *db.Database
}

// New is
func New(e *echo.Echo, c *config.Config, db *db.Database) *App {
	return &App{e, c, db}
}

// SetContext is
func (a *App) SetContext() *App {
	context.New(a.engine)
	return a
}

// SetValidation is
func (a *App) SetValidation() *App {
	validation.New(a.engine)
	return a
}

// SetMiddlewares is
func (a *App) SetMiddlewares() *App {
	middlewares(a.engine)
	return a
}

// SetRoutes is
func (a *App) SetRoutes() *App {
	routes(a.engine, a.config, a.database)
	return a
}

// Run is
func (a *App) Run() {
	c := a.config
	if c.App.Env == "production" {
		a.engine.Logger.Fatal(a.engine.StartTLS(":"+c.App.Port, c.SSL.Cert, c.SSL.Key))
	} else {
		a.engine.Logger.Fatal(a.engine.Start(":" + c.App.Port))
	}
}
