package app

import (
	"go-rest-echo/app/context"
	"go-rest-echo/app/validation"
	"go-rest-echo/config"
	"go-rest-echo/db"
	"go-rest-echo/external"
	"go-rest-echo/service"

	"github.com/labstack/echo/v4"
)

// NewApplicationAndServe is
func NewApplicationAndServe(c *config.Config, db *db.Database, s *service.Service, ex *external.External) {
	e := echo.New()

	context.New(e)
	validation.New(e)
	middlewares(e)
	routes(e, c, db, s, ex)

	e.Logger.Fatal(e.Start(":" + c.App.Port))
}
