package users

import (
	"github.com/labstack/echo/v4"
)

// UserRepository is
type UserRepository interface {
	Fetch() (*[]User, error)
	Get(string) (*User, error)

	Create(*User) error
	Update(*User, string) error
	Delete(string) error
}

// UserUsecase is
type UserUsecase interface {
	Fetch() (*[]User, error)
	Get(string) (*User, error)

	Create(*User) error
	Update(*User, string) error
	Delete(string) error
}

// UserDelivery is
type UserDelivery interface {
	Fetch(echo.Context) error
	Get(echo.Context) error

	Create(echo.Context) error
	Update(echo.Context) error
	Delete(echo.Context) error
}
