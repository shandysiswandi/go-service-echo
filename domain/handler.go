package domain

import (
	"fmt"
	"go-service-echo/app/library/redis"
	"go-service-echo/app/library/sentry"
	"go-service-echo/app/library/token"
	"go-service-echo/app/response"
	"go-service-echo/config/constant"
	"go-service-echo/infrastructure/gormdb"
	"go-service-echo/infrastructure/jsonplaceholder"
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	// DefaultHandler is
	DefaultHandler struct {
		database        *gormdb.Database
		token           *token.Token
		redis           *redis.Redis
		sentry          *sentry.Sentry
		jsonPlaceHolder *jsonplaceholder.JSONPlaceHolder
	}

	// DefaultHandlerConfig is
	DefaultHandlerConfig struct {
		Database        *gormdb.Database
		Token           *token.Token
		Redis           *redis.Redis
		Sentry          *sentry.Sentry
		JSONPlaceHolder *jsonplaceholder.JSONPlaceHolder
	}
)

// NewDefaultHandler is
func NewDefaultHandler(dhc *DefaultHandlerConfig) *DefaultHandler {
	return &DefaultHandler{
		database:        dhc.Database,
		token:           dhc.Token,
		redis:           dhc.Redis,
		sentry:          dhc.Sentry,
		jsonPlaceHolder: dhc.JSONPlaceHolder,
	}
}

// Default is
func (def *DefaultHandler) Default(c echo.Context) error {
	check := c.QueryParam("check")
	msg := ""

	switch check {
	case "db":
		if def.database == nil {
			msg = "checking database -> | gorm (false) | mongo (false)"
		} else {
			msg = fmt.Sprintf("checking database -> | gorm (%v) | mongo (%v)", def.database.SQL != nil, def.database.SQL != nil)
		}
	case "token":
		if def.token == nil {
			msg = "checking token -> | jwt (false) | paseto (false)"
		} else {
			isJWT := false
			isPaseto := false
			if def.token.GetTokenType() == constant.JWT {
				isJWT = true
			}
			if def.token.GetTokenType() == constant.Paseto {
				isPaseto = true
			}
			msg = fmt.Sprintf("checking token -> | jwt (%v) | paseto (%v)", isJWT, isPaseto)
		}
	case "sentry":
		if def.sentry == nil {
			msg = "checking sentry -> (false)"
		} else {
			def.sentry.Message("test sentry")
			msg = fmt.Sprintf("checking sentry -> (%v)", true)
		}
	default:
		msg = "welcome home"
	}

	return response.NewSuccess(c, http.StatusOK, msg, nil)
}

// Favicon is
func (def *DefaultHandler) Favicon(c echo.Context) error {
	return c.File("resource/media/favicon.ico")
}

// CORS is
func (def *DefaultHandler) CORS(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

// ExampleExternalCall is
func (def *DefaultHandler) ExampleExternalCall(c echo.Context) error {
	fetch, err := def.jsonPlaceHolder.FetchPost()
	if err != nil {
		return c.String(502, err.Error())
	}

	get, err := def.jsonPlaceHolder.GetPost(1)
	if err != nil {
		return c.String(502, err.Error())
	}

	pCreate := jsonplaceholder.Post{UserID: 1, ID: 1, Title: "title", Body: "body"}
	create, err := def.jsonPlaceHolder.CreatePost(pCreate)
	if err != nil {
		return c.String(502, err.Error())
	}

	pUpdate := jsonplaceholder.Post{UserID: 1, ID: 1, Title: "title", Body: "body"}
	update, err := def.jsonPlaceHolder.UpdatePost(pUpdate, 1)
	if err != nil {
		return c.String(502, err.Error())
	}

	if err = def.jsonPlaceHolder.DeletePost(1); err != nil {
		return c.String(502, err.Error())
	}

	return response.NewSuccess(c, http.StatusOK, "Welcome to Check Externals", map[string]interface{}{
		"jsonplaceholder": map[string]interface{}{
			"fetch_post":  fetch[1:2],
			"get_post":    get,
			"create_post": create,
			"update_post": update,
			"delete_post": true,
		},
	})
}
