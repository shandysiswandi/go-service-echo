package middlewares

import (
	"fmt"
	"go-service-echo/app/context"
	"go-service-echo/app/library/token"
	"go-service-echo/config/constant"
	"go-service-echo/util/stringy"
	"net/http"

	sentryecho "github.com/getsentry/sentry-go/echo"
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

// Middlewares is
type Middlewares struct {
	engine *echo.Echo
}

// New is
func New(e *echo.Echo) *Middlewares {
	return &Middlewares{e}
}

// PreRouter is middleware to the chain which is run before router.
func (mid *Middlewares) PreRouter() *Middlewares {
	mid.engine.Pre(middleware.RemoveTrailingSlash())

	return mid
}

// PraRouter is middleware to the chain which is run after router.
func (mid *Middlewares) PraRouter(token *token.Token) *Middlewares {
	mid.engine.Use(middleware.Recover())

	mid.engine.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
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
		MaxAge: 2 * 3600,
	}))

	mid.engine.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(cc echo.Context) error {
			ctx := cc.(*context.CustomContext)

			paths := stringy.Split(ctx.Request().URL.Path, "/")
			if len(paths) < 2 {
				return next(ctx)
			}

			if paths[0] != "api" {
				return next(ctx)
			}

			auth := ctx.Request().Header.Get(echo.HeaderAuthorization)
			l := len(constant.AuthScheme)
			if len(auth) > l+1 && auth[:l] == constant.AuthScheme {
				data, err := token.VerifyAccessToken(auth[l+1:])
				if err != nil {
					return ctx.HandleErrors(err)
				}

				ctx.Set("user", data)
				return next(ctx)
			}

			return ctx.Unauthorized(nil)
		}
	})

	mid.engine.Use(middleware.BodyLimit("8M"))
	mid.engine.Use(middleware.GzipWithConfig(middleware.GzipConfig{Level: 5}))
	mid.engine.Use(middleware.Secure())
	mid.engine.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: fmt.Sprintf("%s | %s | %s | %s | %s | %s | %s\n", time, method, uri, ip, status, err, latency),
	}))

	mid.engine.Use(sentryecho.New(sentryecho.Options{
		Repanic:         true,
		WaitForDelivery: false,
		Timeout:         0,
	}))

	return mid
}
