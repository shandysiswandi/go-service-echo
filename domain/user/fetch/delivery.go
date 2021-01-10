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
	u := []entity.User{}

	// usecase
	if err = Usecase(&u); err != nil {
		return c.HandleErrors(err)
	}

	return c.Success(http.StatusOK, "create user", u)
}
