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
	u := entity.User{}
	id := c.Param("id")

	// usecase
	if err = Usecase(&u, id); err != nil {
		return c.HandleErrors(err)
	}

	type H map[string]interface{}
	return c.Success(http.StatusOK, "delete user", H{"deleted_at": u.DeletedAt})
}
