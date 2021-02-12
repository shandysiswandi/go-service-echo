package welcomes

import (
	"go-rest-echo/app/context"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Web is
type Web struct {
	usecase *Usecase
}

// NewWeb is
func NewWeb(u *Usecase) *Web {
	return &Web{u}
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
		"mysql":      w.usecase.CheckDatabaseMysql(),
		"postgresql": w.usecase.CheckDatabasePostgresql(),
		"mongo":      w.usecase.CheckDatabaseMongo(),
	})
}

// CheckLibrary is
func (w *Web) CheckLibrary(cc echo.Context) error {
	c := cc.(*context.CustomContext)

	return c.Success(http.StatusOK, "Welcome to Check Libraries", map[string]interface{}{
		"jwt":    w.usecase.CheckLibraryJWT(),
		"sentry": w.usecase.CheckLibrarySentry(),
		"redis":  w.usecase.CheckLibraryRedis(),
	})
}

// CheckExternal is
func (w *Web) CheckExternal(cc echo.Context) error {
	c := cc.(*context.CustomContext)

	data, err := w.usecase.CheckExternalJSONPlaceHolder()
	if err != nil {
		return c.String(502, err.Error())
	}

	return c.Success(http.StatusOK, "Welcome to Check Externals", map[string]interface{}{
		"jsonplaceholder": data,
	})
}
