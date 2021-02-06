package welcomes

import (
	"go-rest-echo/app/context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type web struct {
	// usecase Usecase
}

// NewWeb is
func NewWeb() WelcomeDelivery {
	return &web{}
}

func (web) Home(cc echo.Context) error {
	// extend echo.Context
	c := cc.(*context.CustomContext)

	// response
	return c.Success(http.StatusOK, "Welcome to our API", nil)
}

func (web) MonitorDatabase(cc echo.Context) error {
	// extend echo.Context
	c := cc.(*context.CustomContext)

	// response
	return c.Success(http.StatusOK, "Welcome to Monitor Databases", nil)
}

func (web) MonitorService(cc echo.Context) error {
	// extend echo.Context
	c := cc.(*context.CustomContext)

	// response
	return c.Success(http.StatusOK, "Welcome to Monitor Services", nil)
}
