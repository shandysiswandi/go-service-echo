package delete

import (
	"go-rest-echo/entity"
	"go-rest-echo/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Delivery is
func Delivery(cc echo.Context) (err error) {
	c := cc.(*helper.Context)
	u := entity.User{}
	id := c.Param("id")

	// usecase
	if err = Usecase(&u, id); err != nil {
		return c.HandleErrors(err)
	}

	return c.Success(http.StatusOK, "create user", u)
}
