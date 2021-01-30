package authentication

import (
	"go-rest-echo/app/context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type web struct {
	usecase AuthUsecase
}

// NewWeb is
func NewWeb(u AuthUsecase) AuthDelivery {
	return &web{usecase: u}
}

func (w *web) Login(cc echo.Context) (err error) {
	// extend echo.Context
	c := cc.(*context.CustomContext)

	// define variables
	pl := new(PayloadLogin)

	// binding
	if err = c.Bind(pl); err != nil {
		return c.BadRequest(err)
	}

	// validation
	if err = c.Validate(pl); err != nil {
		return c.UnprocessableEntity(err)
	}

	// usecase
	result, err := w.usecase.Login(pl)
	if err != nil {
		log.Println(err)
		return c.HandleErrors(err)
	}

	// response
	return c.Success(http.StatusOK, "auth login", result)
}
