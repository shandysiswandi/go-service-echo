package fetch

import (
	"go-rest-echo/entity"
	"go-rest-echo/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Delivery is
func Delivery(c echo.Context) (err error) {
	u := []entity.User{}

	// usecase
	if err = Usecase(&u); err != nil {
		return response.HandleErrors(c, err)
	}

	return response.Success(c, http.StatusCreated, "create user", u)
}
