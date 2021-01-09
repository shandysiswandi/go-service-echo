package delete

import (
	"go-rest-echo/entity"
	"go-rest-echo/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Delivery is
func Delivery(c echo.Context) (err error) {
	u := entity.User{}
	id := c.Param("id")

	// usecase
	if err = Usecase(&u, id); err != nil {
		return response.HandleErrors(c, err)
	}

	return response.Success(c, http.StatusOK, "create user", u)
}
