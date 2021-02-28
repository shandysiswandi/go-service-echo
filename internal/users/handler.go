package users

import (
	"go-service-echo/app/context"
	"net/http"

	"github.com/labstack/echo/v4"
)

// UserHandler is
type UserHandler struct {
	usecase *UserUsecase
}

// NewDelivery is
func NewDelivery(u *UserUsecase) *UserHandler {
	return &UserHandler{usecase: u}
}

// Fetch is
func (d *UserHandler) Fetch(cc echo.Context) (err error) {
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

// Get is
func (d *UserHandler) Get(cc echo.Context) (err error) {
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

// Create is
func (d *UserHandler) Create(cc echo.Context) (err error) {
	// extend echo.Context
	c := cc.(*context.CustomContext)

	// define variables
	u := new(User)

	// binding
	if err = c.Bind(u); err != nil {
		return c.BadRequest(err)
	}

	// validation
	if err = c.Validate(u); err != nil {
		return c.UnprocessableEntity(err)
	}

	// usecase
	if err = d.usecase.Create(u); err != nil {
		return c.HandleErrors(err)
	}

	// response
	return c.Success(http.StatusCreated, "create user", u)
}

// Update is
func (d *UserHandler) Update(cc echo.Context) (err error) {
	// extend echo.Context
	c := cc.(*context.CustomContext)

	// define variables
	u := new(User)
	id := c.Param("id")

	// binding
	if err = c.Bind(u); err != nil {
		return c.BadRequest(err)
	}

	// validation
	if err = c.Validate(u); err != nil {
		return c.UnprocessableEntity(err)
	}

	// usecase
	if err = d.usecase.Update(u, id); err != nil {
		return c.HandleErrors(err)
	}

	// response
	return c.Success(http.StatusOK, "update task", u)
}

// Delete is
func (d *UserHandler) Delete(cc echo.Context) (err error) {
	// extend echo.Context
	c := cc.(*context.CustomContext)

	// define variables
	u := User{}
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
