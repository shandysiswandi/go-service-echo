package internal

import (
	"fmt"
	"go-service-echo/app/context"
	cc "go-service-echo/app/library/redis"
	dd "go-service-echo/app/library/sentry"
	bb "go-service-echo/app/library/token"
	"go-service-echo/config/constant"
	aa "go-service-echo/infrastructure/database"
	ee "go-service-echo/infrastructure/jsonplaceholder"
	"net/http"

	"github.com/labstack/echo/v4"
)

// DefaultInternal is
type DefaultInternal struct {
	database        *aa.Database
	token           *bb.Token
	redis           *cc.Redis
	sentry          *dd.Sentry
	jsonPlaceHolder *ee.JSONPlaceHolder
}

// NewHandler is
func NewHandler(a *aa.Database, b *bb.Token, c *cc.Redis, d *dd.Sentry, e *ee.JSONPlaceHolder) *DefaultInternal {
	return &DefaultInternal{
		database:        a,
		token:           b,
		redis:           c,
		sentry:          d,
		jsonPlaceHolder: e,
	}
}

// Default is
func (def *DefaultInternal) Default(ctx echo.Context) error {
	c := ctx.(*context.CustomContext)
	check := c.QueryParam("check")
	msg := ""

	switch check {
	case "db":
		if def.database == nil {
			msg = "checking database -> | gorm (false) | mongo (false)"
		} else {
			msg = fmt.Sprintf("checking database -> | gorm (%v) | mongo (%v)", def.database.SQL != nil, def.database.Mongo != nil)
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

	return c.Success(http.StatusOK, msg, nil)
}

// Favicon is
func (def *DefaultInternal) Favicon(c echo.Context) error {
	return c.File("resource/media/favicon.ico")
}

// CORS is
func (def *DefaultInternal) CORS(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

// ExampleExternalCall is
func (def *DefaultInternal) ExampleExternalCall(cc echo.Context) error {
	c := cc.(*context.CustomContext)

	fetch, err := def.jsonPlaceHolder.FetchPost()
	if err != nil {
		return c.String(502, err.Error())
	}

	get, err := def.jsonPlaceHolder.GetPost(1)
	if err != nil {
		return c.String(502, err.Error())
	}

	pCreate := ee.Post{UserID: 1, ID: 1, Title: "title", Body: "body"}
	create, err := def.jsonPlaceHolder.CreatePost(pCreate)
	if err != nil {
		return c.String(502, err.Error())
	}

	pUpdate := ee.Post{UserID: 1, ID: 1, Title: "title", Body: "body"}
	update, err := def.jsonPlaceHolder.UpdatePost(pUpdate, 1)
	if err != nil {
		return c.String(502, err.Error())
	}

	if err = def.jsonPlaceHolder.DeletePost(1); err != nil {
		return c.String(502, err.Error())
	}

	return c.Success(http.StatusOK, "Welcome to Check Externals", map[string]interface{}{
		"jsonplaceholder": map[string]interface{}{
			"fetch_post":  fetch[1:2],
			"get_post":    get,
			"create_post": create,
			"update_post": update,
			"delete_post": true,
		},
	})
}
