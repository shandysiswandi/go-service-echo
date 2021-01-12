package app

import (
	"go-rest-echo/app/context"
	"go-rest-echo/app/middleware"
	"go-rest-echo/app/route"
	"go-rest-echo/app/validation"
	"go-rest-echo/helper"

	"github.com/labstack/echo/v4"
)

type app struct{}

// Interface is
type Interface interface {
	Start()
}

// NewApplication is
func NewApplication() Interface {
	return &app{}
}

// Start is
func (app) Start() {
	e := echo.New()

	context.Initialize(e)
	validation.Initialize(e)
	middleware.Initialize(e)
	route.Initialize(e)

	e.Logger.Fatal(e.Start(":" + helper.Env("PORT")))
}
