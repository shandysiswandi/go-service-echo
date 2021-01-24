package tasks

import (
	"go-rest-echo/app/context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type delivery struct {
	usecase TaskUsecase
}

// NewDelivery is
func NewDelivery(u TaskUsecase) TaskDelivery {
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
	return c.Success(http.StatusOK, "fetch tasks", result)
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
	return c.Success(http.StatusOK, "get task", result)
}

func (d *delivery) Create(cc echo.Context) (err error) {
	// extend echo.Context
	c := cc.(*context.CustomContext)

	// define variables
	t := Task{}

	// binding
	if err = c.Bind(&t); err != nil {
		return c.BadRequest(err)
	}

	// validation
	if err = c.Validate(t); err != nil {
		return c.UnprocessableEntity(err)
	}

	// usecase
	if _, err = d.usecase.Create(&t); err != nil {
		return c.HandleErrors(err)
	}

	// response
	return c.Success(http.StatusCreated, "create task", t)
}

func (d *delivery) Update(cc echo.Context) (err error) {
	// extend echo.Context
	c := cc.(*context.CustomContext)

	// define variables
	t := Task{}
	id := c.Param("id")

	// binding
	if err = c.Bind(&t); err != nil {
		return c.BadRequest(err)
	}

	// validation
	if err = c.Validate(t); err != nil {
		return c.UnprocessableEntity(err)
	}

	// usecase
	if err = d.usecase.Update(&t, id); err != nil {
		return c.HandleErrors(err)
	}

	// response
	return c.Success(http.StatusOK, "update task", t)
}

func (d *delivery) Delete(cc echo.Context) (err error) {
	// extend echo.Context
	c := cc.(*context.CustomContext)

	// define variables
	t := Task{}
	id := c.Param("id")

	// usecase
	if err = d.usecase.Delete(id); err != nil {
		return c.HandleErrors(err)
	}

	// response
	return c.Success(http.StatusOK, "delete user", map[string]interface{}{
		"deleted_at": t.DeletedAt,
	})
}
