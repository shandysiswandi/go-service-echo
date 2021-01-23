package task

import (
	"go-rest-echo/app/context"
	"go-rest-echo/entity"
	"go-rest-echo/usecase/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type delivery struct {
	usecase user.Usecase
}

// NewDelivery is
func NewDelivery(u user.Usecase) user.Delivery {
	return &delivery{usecase: u}
}

func (d *delivery) Fetch(cc echo.Context) (err error) {
	// extend echo.Context
	c := cc.(*context.CustomContext)

	// usecase
	result, err := d.usecase.Fetch()
	if err != nil {
		return c.HandleErrors(err)
	}

	// response
	return c.Success(http.StatusOK, "fetch users", result)
}

func (d *delivery) Get(cc echo.Context) (err error) {
	// extend echo.Context
	c := cc.(*context.CustomContext)

	// define variables
	id := c.Param("id")

	// usecase
	result, err := d.usecase.Get(id)
	if err != nil {
		return c.HandleErrors(err)
	}

	// response
	return c.Success(http.StatusOK, "get user", result)
}

func (d *delivery) Create(cc echo.Context) (err error) {
	// extend echo.Context
	c := cc.(*context.CustomContext)

	// define variables
	u := entity.User{}

	// binding
	if err = c.Bind(&u); err != nil {
		return c.BadRequest(err)
	}

	// validation
	if err = c.Validate(u); err != nil {
		return c.UnprocessableEntity(err)
	}

	// usecase
	if err = d.usecase.Create(&u); err != nil {
		return c.HandleErrors(err)
	}

	// response
	return c.Success(http.StatusCreated, "create user", u)
}

func (d *delivery) Update(cc echo.Context) (err error) {
	// extend echo.Context
	c := cc.(*context.CustomContext)

	// define variables
	u := entity.User{}
	id := c.Param("id")

	// binding
	if err = c.Bind(&u); err != nil {
		return c.BadRequest(err)
	}

	// validation
	if err = c.Validate(u); err != nil {
		return c.UnprocessableEntity(err)
	}

	// usecase
	if err = d.usecase.Update(&u, id); err != nil {
		return c.HandleErrors(err)
	}

	// response
	return c.Success(http.StatusOK, "update task", u)
}

func (d *delivery) Delete(cc echo.Context) (err error) {
	// extend echo.Context
	c := cc.(*context.CustomContext)

	// define variables
	u := entity.User{}
	id := c.Param("id")

	// usecase
	if err = d.usecase.Delete(id); err != nil {
		return c.HandleErrors(err)
	}

	// response
	return c.Success(http.StatusOK, "delete user", map[string]interface{}{
		"deleted_at": u.DeletedAt,
	})
}
