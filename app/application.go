package app

import (
	"go-rest-echo/app/context"
	"go-rest-echo/app/middleware"
	"go-rest-echo/app/validation"
	"go-rest-echo/config"
	"go-rest-echo/db"

	"github.com/labstack/echo/v4"
)

// NewApplicationAndServe is
func NewApplicationAndServe(conf *config.Config, db *db.Database) {
	e := echo.New()

	// extend echo Context
	context.New(e)

	// register validation
	validation.New(e)

	// middlewares
	middleware.RemoveTrailingSlash(e)
	middleware.Logger(e)
	middleware.Recover(e)
	middleware.Cors(e)
	middleware.BodyLimit(e)
	middleware.Gzip(e)
	middleware.Secure(e)

	routeWithoutJwt(e, conf)
	routeWithJwt(e, conf, db)

	e.Logger.Fatal(e.Start(":" + conf.App.Port))
}
