package app

import (
	"fmt"
	"go-rest-echo/app/context"
	"go-rest-echo/app/validation"
	"go-rest-echo/domain/task"
	"go-rest-echo/domain/user"
	"go-rest-echo/helper"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	time    = "\033[33mTIME\033[0m: ${time_rfc3339}"
	method  = "\033[33mMETHOD\033[0m: ${method}"
	uri     = "\033[33mURI\033[0m: ${uri}"
	ip      = "\033[33mIP\033[0m: ${remote_ip}"
	status  = "\033[33mSTATUS\033[0m: ${status}"
	err     = "\033[33mERROR\033[0m: ${error}"
	latency = "\033[33mLATENCY\033[0m: ${latency_human}"
)

type app struct{}

// Interface is
type Interface interface {
	Start()
	initializeMiddlewares(*echo.Echo)
	initializeRoutes(*echo.Echo)
}

// NewApplication is
func NewApplication() Interface {
	return &app{}
}

// Start is
func (a app) Start() {
	e := echo.New()

	context.NewContext(e)
	validation.NewValidation(e)

	a.initializeMiddlewares(e)
	a.initializeRoutes(e)

	e.Logger.Fatal(e.Start(":" + helper.Env("PORT")))
}

func (app) initializeMiddlewares(e *echo.Echo) {
	// Remove Trailing Slash
	e.Pre(middleware.RemoveTrailingSlash())
	// logger
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: fmt.Sprintf("%s | %s | %s | %s | %s | %s | %s\n", time, method, uri, ip, status, err, latency),
	}))
	// recover panic
	e.Use(middleware.Recover())
	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{
			echo.HeaderAccept,
			echo.HeaderAcceptEncoding,
			echo.HeaderAccessControlAllowHeaders,
			echo.HeaderAuthorization,
			echo.HeaderContentType,
			echo.HeaderContentLength,
			echo.HeaderOrigin,
			echo.HeaderXCSRFToken,
		},
		MaxAge: 86400,
	}))
	// body limit
	e.Use(middleware.BodyLimit("8M"))
	// Gzip compress
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{Level: 5}))
	// secure
	e.Use(middleware.Secure())
}

func (app) initializeRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("%s ----> %s", helper.Env("NAME"), helper.Env("ENV")))
	})

	task.NewRouter().Initialize(e, "/tasks")
	user.NewRouter().Initialize(e, "/users")
}
