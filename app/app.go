package app

import (
	"go-service-echo/app/context"
	"go-service-echo/app/validation"
	"go-service-echo/config"
	"go-service-echo/db"

	"github.com/labstack/echo/v4"
)

// New is
func New(c *config.Config, db *db.Database) {
	e := echo.New()

	context.New(e)
	validation.New(e)

	middlewares(e)
	routes(e, c, db)

	e.Logger.Fatal(e.Start(":" + c.App.Port))
}
