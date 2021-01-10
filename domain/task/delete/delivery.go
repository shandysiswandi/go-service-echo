package delete

import (
	"go-rest-echo/app/context"
	"go-rest-echo/entity"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Delivery is
func Delivery(cc echo.Context) (err error) {
	c := cc.(*context.CustomContext)
	t := entity.Task{}
	id := c.Param("id")

	// usecase
	if err = Usecase(&t, id); err != nil {
		return c.HandleErrors(err)
	}

	return c.Success(http.StatusOK, "delete task", t)
}
