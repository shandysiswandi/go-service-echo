package fetch

import (
	"go-rest-echo/app/context"
	"go-rest-echo/entity"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Delivery is
func Delivery(cc echo.Context) (err error) {
	c := cc.(*context.CustomContext)
	t := &[]entity.Task{}

	// usecase
	if err = Usecase(t); err != nil {
		return c.HandleErrors(err)
	}

	return c.Success(http.StatusOK, "fetch users", t)
}
