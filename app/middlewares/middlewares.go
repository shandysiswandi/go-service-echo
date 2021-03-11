package middlewares

import (
	"fmt"
	"go-service-echo/app/library/token"
	"go-service-echo/app/response"
	"go-service-echo/config/constant"
	"go-service-echo/util/arrays"
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
	token  *token.Token
}

// New is
func New(e *echo.Echo, token *token.Token) *Middlewares {
	/* ----------------------------------------------------------- */
	/* Official Echo Middleware
	/*
	/* Pre is middleware to the chain which is run before router.
	/* ----------------------------------------------------------- */
	e.Pre(middleware.RemoveTrailingSlash())

	/* ----------------------------------------------------------- */
	/* Official Echo Middleware
	/*
	/* Pra is middleware to the chain which is run after router.
	/* ----------------------------------------------------------- */
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: fmt.Sprintf("%s | %s | %s | %s | %s | %s | %s\n", time, method, uri, ip, status, err, latency),
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Use(middleware.BodyLimit("8M"))
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{Level: 5}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
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

	/* ----------------------------------------------------------- */
	/* Third-Party Echo Middleware
	/*
	/* ----------------------------------------------------------- */
	e.Use(sentryecho.New(sentryecho.Options{
		Repanic:         true,
		WaitForDelivery: false,
		Timeout:         0,
	}))

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {

			paths := arrays.Split(ctx.Request().URL.Path, "/")
			if len(paths) < 2 {
				return next(ctx)
			}

			if paths[0] != constant.AuthRoutePrefix {
				return next(ctx)
			}

			auth := ctx.Request().Header.Get(echo.HeaderAuthorization)
			length := len(constant.AuthScheme)
			if len(auth) > length+1 && auth[:length] == constant.AuthScheme {
				data, err := token.VerifyAccessToken(auth[length+1:])
				if err != nil {
					return response.HandleErrors(ctx, err)
				}

				ctx.Set("user", data)
				return next(ctx)
			}

			return response.Unauthorized(ctx, nil)
		}
	})

	/* return nil */
	return nil
}
