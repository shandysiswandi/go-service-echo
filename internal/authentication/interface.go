package authentication

import (
	"github.com/labstack/echo/v4"
)

// AuthDelivery is
type AuthDelivery interface {
	Login(echo.Context) error
}

// AuthUsecase is
type AuthUsecase interface {
	Login(*PayloadLogin) (*ResponseLogin, error)
}
