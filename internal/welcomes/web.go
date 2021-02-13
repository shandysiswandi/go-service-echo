package welcomes

import (
	"go-rest-echo/app/context"
	"go-rest-echo/app/library/jwtlib"
	"go-rest-echo/app/library/redislib"
	"go-rest-echo/app/library/sentrylib"
	"go-rest-echo/db"
	"go-rest-echo/external/jsonplaceholder"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Web is
type Web struct {
	database        *db.Database
	jwt             *jwtlib.JWT
	redis           *redislib.Redis
	sentry          *sentrylib.Sentry
	jsonPlaceHolder *jsonplaceholder.JSONPlaceHolder
}

// NewWeb is
func NewWeb(db *db.Database, j *jwtlib.JWT, r *redislib.Redis, s *sentrylib.Sentry, jph *jsonplaceholder.JSONPlaceHolder) *Web {
	return &Web{db, j, r, s, jph}
}

// Home is
func (w *Web) Home(cc echo.Context) error {
	c := cc.(*context.CustomContext)

	return c.Success(http.StatusOK, "Welcome to our API", map[string]interface{}{
		"route_check_database": "/check-database",
		"route_check_library":  "/check-library",
		"route_check_external": "/check-external",
	})
}

// CheckDatabase is
func (w *Web) CheckDatabase(cc echo.Context) error {
	c := cc.(*context.CustomContext)

	return c.Success(http.StatusOK, "Welcome to Check Databases", map[string]interface{}{
		"mysql_or_postgresql": w.database.SQL != nil,
		"mongo":               w.database.Mongo != nil,
	})
}

// CheckLibrary is
func (w *Web) CheckLibrary(cc echo.Context) error {
	c := cc.(*context.CustomContext)

	return c.Success(http.StatusOK, "Welcome to Check Libraries", map[string]interface{}{
		"jwt":    w.jwt != nil,
		"sentry": w.sentry != nil,
		"redis":  w.redis != nil,
	})
}

// CheckExternal is
func (w *Web) CheckExternal(cc echo.Context) error {
	c := cc.(*context.CustomContext)

	fetch, err := w.jsonPlaceHolder.FetchPost()
	if err != nil {
		return c.String(502, err.Error())
	}

	get, err := w.jsonPlaceHolder.GetPost(1)
	if err != nil {
		return c.String(502, err.Error())
	}

	pCreate := jsonplaceholder.Post{UserID: 1, ID: 1, Title: "title", Body: "body"}
	create, err := w.jsonPlaceHolder.CreatePost(pCreate)
	if err != nil {
		return c.String(502, err.Error())
	}

	pUpdate := jsonplaceholder.Post{UserID: 1, ID: 1, Title: "title", Body: "body"}
	update, err := w.jsonPlaceHolder.UpdatePost(pUpdate, 1)
	if err != nil {
		return c.String(502, err.Error())
	}

	if err = w.jsonPlaceHolder.DeletePost(1); err != nil {
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
