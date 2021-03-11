package users

import (
	"go-service-echo/app/response"
	"go-service-echo/app/validation"
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

func (d *userHandler) Fetch(c echo.Context) (err error) {
	// usecase
	result, err := d.usecase.Fetch()
	if err != nil {
		return response.HandleErrors(c, err)
	}

	// response
	return response.NewSuccess(c, http.StatusOK, "fetch users", result.Transform())
}

func (d *userHandler) Get(c echo.Context) (err error) {
	// define variables
	id := c.Param("id")

	// usecase
	result, err := d.usecase.Get(id)
	if err != nil {
		return response.HandleErrors(c, err)
	}

	// response
	return response.NewSuccess(c, http.StatusOK, "get user", result)
}

func (d *userHandler) GetByEmail(c echo.Context) (err error) {
	// define variables
	email := c.Param("email")

	// validation
	if err := validation.ValidateVar(email, "required,email,min:5"); err != nil {
		return response.UnprocessableEntityVar(c, err)
	}

	// usecase
	result, err := d.usecase.GetByEmail(email)
	if err != nil {
		return response.HandleErrors(c, err)
	}

	// response
	return response.NewSuccess(c, http.StatusOK, "get user by email", result)
}

func (d *userHandler) Create(c echo.Context) (err error) {
	// define variables
	u := new(UserCreatePayload)

	// binding
	if err = c.Bind(u); err != nil {
		return response.BadRequest(c, err)
	}

	// validation
	if err = c.Validate(u); err != nil {
		return response.UnprocessableEntity(c, err)
	}

	// usecase
	if err = d.usecase.Create(u); err != nil {
		return response.HandleErrors(c, err)
	}

	// response
	return response.NewSuccess(c, http.StatusCreated, "create user", u)
}

func (d *userHandler) Update(c echo.Context) (err error) {
	// define variables
	u := new(UserUpdatePayload)
	id := c.Param("id")

	// binding
	if err = c.Bind(u); err != nil {
		return response.BadRequest(c, err)
	}

	// validation
	if err = c.Validate(u); err != nil {
		return response.UnprocessableEntity(c, err)
	}

	// usecase
	if err = d.usecase.Update(u, id); err != nil {
		return response.HandleErrors(c, err)
	}

	// response
	return response.NewSuccess(c, http.StatusOK, "update task", u)
}

func (d *userHandler) Delete(c echo.Context) (err error) {
	// define variables
	id := c.Param("id")

	// usecase
	if err = d.usecase.Delete(id); err != nil {
		return response.HandleErrors(c, err)
	}

	// response
	return response.NewSuccess(c, http.StatusOK, "delete user", map[string]interface{}{
		"deleted_at": time.Now(),
	})
}
