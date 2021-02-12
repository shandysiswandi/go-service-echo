package app

import (
	"go-rest-echo/app/context"
	"go-rest-echo/app/validation"
	"go-rest-echo/config"
	"go-rest-echo/db"

	"github.com/labstack/echo/v4"
)

// NewApplicationAndServe is
func NewApplicationAndServe(c *config.Config, db *db.Database) {
	e := echo.New()

	context.New(e)
	validation.New(e)

	middlewares(e)
	routes(e, c, db)

	e.Logger.Fatal(e.Start(":" + c.App.Port))
}
