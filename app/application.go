package app

import (
	"go-rest-echo/app/context"
	"go-rest-echo/app/middleware"
	"go-rest-echo/app/route"
	"go-rest-echo/app/validation"
	"go-rest-echo/config"
	"go-rest-echo/db"

	"github.com/labstack/echo/v4"
)

// NewApplication is
func NewApplication(conf *config.Config, db *db.Database) error {
	// instance of echo framework
	e := echo.New()

	// extend echo Context
	context.NewCustomContext(e)

	// register validation
	validation.NewValidation(e)

	// middlewares
	middleware.RemoveTrailingSlash(e)
	middleware.Logger(e)
	middleware.Recover(e)
	middleware.Cors(e)
	middleware.BodyLimit(e)
	middleware.Gzip(e)
	middleware.Secure(e)

	// routes
	route.HomeRoute(e, conf, db)
	route.TaskRoute(e, conf, db)
	route.UserRoute(e, conf, db)

	// run application
	return e.Start(":" + conf.App.Port)
}
