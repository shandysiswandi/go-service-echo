package user

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
	u := []entity.User{}

	// usecase
	if err = d.usecase.Fetch(&u); err != nil {
		return c.HandleErrors(err)
	}

	return c.Success(http.StatusOK, "create user", u)
}

func (d *delivery) Get(cc echo.Context) (err error) {
	c := cc.(*context.CustomContext)
	u := entity.User{}
	id := c.Param("id")

	// usecase
	if err = d.usecase.Get(&u, id); err != nil {
		return c.HandleErrors(err)
	}

	return c.Success(http.StatusOK, "get user", u)
}

func (d *delivery) Create(cc echo.Context) (err error) {
	c := cc.(*context.CustomContext)
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
	if err = d.usecase.Create(&u); err != nil {
		return c.HandleErrors(err)
	}

	return c.Success(http.StatusCreated, "create user", u)
}

func (d *delivery) Update(cc echo.Context) (err error) {
	c := cc.(*context.CustomContext)
	u := entity.User{}
	id := c.Param("id")

	// binding
	if err = c.Bind(&u); err != nil {
		return c.BadRequest(err)
	}

	// validation
	if err = c.Validate(u); err != nil {
		return c.ValidationFail(err)
	}

	// usecase
	if err = d.usecase.Update(&u, id); err != nil {
		return c.HandleErrors(err)
	}

	return c.Success(http.StatusOK, "get user", u)
}

func (d *delivery) Delete(cc echo.Context) (err error) {
	c := cc.(*context.CustomContext)
	u := entity.User{}
	id := c.Param("id")

	// usecase
	if err = d.usecase.Delete(&u, id); err != nil {
		return c.HandleErrors(err)
	}

	return c.Success(http.StatusOK, "delete user", map[string]interface{}{
		"deleted_at": u.DeletedAt,
	})
}
