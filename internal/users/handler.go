package users

import (
	"go-service-echo/app/context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	usecase UserUsecase
}

// NewUserHandler is
func NewUserHandler(u UserUsecase) UserHandler {
	return &userHandler{u}
}

func (d *userHandler) Fetch(cc echo.Context) (err error) {
	c := cc.(*context.CustomContext)

	// usecase
	result, err := d.usecase.Fetch()
	if err != nil {
		return c.HandleErrors(err)
	}

	// response
	return c.Success(http.StatusOK, "fetch users", result.Transform())
}

func (d *userHandler) Get(cc echo.Context) (err error) {
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

func (d *userHandler) GetByEmail(cc echo.Context) (err error) {
	c := cc.(*context.CustomContext)

	// define variables
	email := c.Param("email")

	// validation
	if err := c.ValidateVar(email, "required,email,min:5"); err != nil {
		return c.UnprocessableEntityVar(err)
	}

	// usecase
	result, err := d.usecase.GetByEmail(email)
	if err != nil {
		return c.HandleErrors(err)
	}

	// response
	return c.Success(http.StatusOK, "get user by email", result)
}

func (d *userHandler) Create(cc echo.Context) (err error) {
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

func (d *userHandler) Update(cc echo.Context) (err error) {
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

func (d *userHandler) Delete(cc echo.Context) (err error) {
	c := cc.(*context.CustomContext)

	// define variables
	id := c.Param("id")

	// usecase
	if err = d.usecase.Delete(id); err != nil {
		return c.HandleErrors(err)
	}

	// response
	return c.Success(http.StatusOK, "delete user", map[string]interface{}{
		"deleted_at": time.Now(),
	})
}
