package create

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

	// binding
	if err = c.Bind(&u); err != nil {
		return c.BadRequest(err)
	}

	// validation
	if err = c.Validate(u); err != nil {
		return c.ValidationFail(err)
	}

	// usecase
	if err = Usecase(&u); err != nil {
		return c.HandleErrors(err)
	}

	return c.Success(http.StatusCreated, "create user", u)
}
