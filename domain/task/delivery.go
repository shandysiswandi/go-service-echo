package task

import (
	"go-rest-echo/app/context"
	"go-rest-echo/entity"
	"net/http"

	"github.com/labstack/echo/v4"
)

type delivery struct {
	usecase Usecase
}

// Delivery is
type Delivery interface {
	Fetch(echo.Context) error
	Get(echo.Context) error
	Create(echo.Context) error
	Update(echo.Context) error
	Delete(echo.Context) error
}

// NewDelivery is
func NewDelivery() Delivery {
	return &delivery{usecase: NewUsecase()}
}

func (d *delivery) Fetch(cc echo.Context) (err error) {
	c := cc.(*context.CustomContext)
	t := []entity.Task{}

	// usecase
	if err = d.usecase.Fetch(&t); err != nil {
		return c.HandleErrors(err)
	}

	return c.Success(http.StatusOK, "fetch tasks", t)
}

func (d *delivery) Get(cc echo.Context) (err error) {
	c := cc.(*context.CustomContext)
	t := entity.Task{}
	id := c.Param("id")

	// usecase
	if err = d.usecase.Get(&t, id); err != nil {
		return c.HandleErrors(err)
	}

	return c.Success(http.StatusOK, "get user", t)
}

func (d *delivery) Create(cc echo.Context) (err error) {
	c := cc.(*context.CustomContext)
	t := entity.Task{}

	// binding
	if err = c.Bind(&t); err != nil {
		return c.BadRequest(err)
	}

	// validation
	if err = c.Validate(t); err != nil {
		return c.ValidationFail(err)
	}

	// usecase
	if err = d.usecase.Create(&t); err != nil {
		return c.HandleErrors(err)
	}

	return c.Success(http.StatusCreated, "create user", t)
}

func (d *delivery) Update(cc echo.Context) (err error) {
	c := cc.(*context.CustomContext)
	t := entity.Task{}
	id := c.Param("id")

	// binding
	if err = c.Bind(&t); err != nil {
		return c.BadRequest(err)
	}

	// validation
	if err = c.Validate(t); err != nil {
		return c.ValidationFail(err)
	}

	// usecase
	if err = d.usecase.Update(&t, id); err != nil {
		return c.HandleErrors(err)
	}

	return c.Success(http.StatusOK, "update task", t)
}

func (d *delivery) Delete(cc echo.Context) (err error) {
	c := cc.(*context.CustomContext)
	t := entity.Task{}
	id := c.Param("id")

	// usecase
	if err = d.usecase.Delete(&t, id); err != nil {
		return c.HandleErrors(err)
	}

	return c.Success(http.StatusOK, "delete user", map[string]interface{}{
		"deleted_at": t.DeletedAt,
	})
}
