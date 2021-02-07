package blogs

import (
	"go-rest-echo/app/context"
	"net/http"

	"github.com/labstack/echo/v4"
)

type web struct {
	usecase Usecase
}

// NewWeb is
func NewWeb(bu Usecase) BlogDelivery {
	return &web{usecase: bu}
}

func (w *web) Fetch(cc echo.Context) (err error) {
	// extend echo.Context
	c := cc.(*context.CustomContext)

	// usecase
	result, err := w.usecase.Fetch()
	if err != nil {
		return c.HandleErrors(err)
	}

	// response
	return c.Success(http.StatusOK, "fetch tasks", result)
}

func (w *web) Get(cc echo.Context) (err error) {
	// extend echo.Context
	c := cc.(*context.CustomContext)

	// define variables
	id := c.Param("id")

	// usecase
	result, err := w.usecase.Get(id)
	if err != nil {
		return c.HandleErrors(err)
	}

	// response
	return c.Success(http.StatusOK, "get task", result)
}

func (w *web) Create(cc echo.Context) (err error) {
	// extend echo.Context
	c := cc.(*context.CustomContext)

	// define variables
	b := BlogPayloadCreate{}

	// binding
	if err = c.Bind(&b); err != nil {
		return c.BadRequest(err)
	}

	// validation
	if err = c.Validate(&b); err != nil {
		return c.UnprocessableEntity(err)
	}

	// usecase
	if err = w.usecase.Create(b); err != nil {
		return c.HandleErrors(err)
	}

	// response
	return c.Success(http.StatusCreated, "create task", b)
}

func (w *web) Update(cc echo.Context) (err error) {
	// extend echo.Context
	c := cc.(*context.CustomContext)

	// define variables
	b := new(Blog)
	id := c.Param("id")

	// binding
	if err = c.Bind(b); err != nil {
		return c.BadRequest(err)
	}

	// validation
	if err = c.Validate(b); err != nil {
		return c.UnprocessableEntity(err)
	}

	// usecase
	if err = w.usecase.Update(b, id); err != nil {
		return c.HandleErrors(err)
	}

	// response
	return c.Success(http.StatusOK, "update task", b)
}

func (w *web) Delete(cc echo.Context) (err error) {
	// extend echo.Context
	c := cc.(*context.CustomContext)

	// define variables
	// b := new(Blog)
	id := c.Param("id")

	// usecase
	if err = w.usecase.Delete(id); err != nil {
		return c.HandleErrors(err)
	}

	// response
	return c.Success(http.StatusOK, "delete user", map[string]interface{}{
		"is_deleted": true,
	})
}
