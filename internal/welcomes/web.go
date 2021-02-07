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
	// extend echo.Context
	c := cc.(*context.CustomContext)

	// response
	return c.Success(http.StatusOK, "Welcome to our API", nil)
}

// MonitorDatabase is
func (w *Web) MonitorDatabase(cc echo.Context) error {
	// extend echo.Context
	c := cc.(*context.CustomContext)

	// usecases
	mysql := w.usecase.CheckDatabaseMysql()
	postgresql := w.usecase.CheckDatabasePostgresql()
	mongo := w.usecase.CheckDatabaseMongo()

	// response
	return c.Success(http.StatusOK, "Welcome to Monitor Databases", map[string]interface{}{
		"mysql":      mysql,
		"postgresql": postgresql,
		"mongo":      mongo,
	})
}

// MonitorService is
func (w *Web) MonitorService(cc echo.Context) error {
	// extend echo.Context
	c := cc.(*context.CustomContext)

	// usecases
	jwt := w.usecase.CheckServiceJWT()
	sentry := w.usecase.CheckServiceSentry()
	redis := w.usecase.CheckServiceRedis()

	// response
	return c.Success(http.StatusOK, "Welcome to Monitor Services", map[string]interface{}{
		"jwt":    jwt,
		"sentry": sentry,
		"redis":  redis,
	})
}

// MonitorExternal is
func (w *Web) MonitorExternal(cc echo.Context) error {
	// extend echo.Context
	c := cc.(*context.CustomContext)

	// usecases
	data, err := w.usecase.CheckExternalJSONPlaceHolder()
	if err != nil {
		return c.String(502, err.Error())
	}

	// response
	return c.Success(http.StatusOK, "Welcome to Monitor Externals", map[string]interface{}{
		"jsonplaceholder": data,
	})
}
