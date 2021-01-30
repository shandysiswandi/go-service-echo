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

	// extend echo Context
	context.New(e)

	// register validation
	validation.New(e)

	middlewares(e)

	routeWithoutJwt(e, c, db)
	routeWithJwt(e, c, db)

	e.Logger.Fatal(e.Start(":" + c.App.Port))
}
