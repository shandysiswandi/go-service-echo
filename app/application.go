package app

import (
	"os"

	"go-rest-echo/app/context"
	"go-rest-echo/app/middleware"
	"go-rest-echo/app/route"
	"go-rest-echo/app/validation"

	"github.com/labstack/echo/v4"
)

// NewApplication is
func NewApplication() {
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
	route.TaskRoute(e)
	route.UserRoute(e)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
