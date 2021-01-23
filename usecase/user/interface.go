package user

import (
	"go-rest-echo/entity"

	"github.com/labstack/echo/v4"
)

// Repository is
type Repository interface {
	Fetch() ([]*entity.User, error)
	Get(string) (*entity.User, error)

	Create(*entity.User) error
	Update(*entity.User, string) error
	Delete(string) error
}

// Usecase is
type Usecase interface {
	Fetch() ([]*entity.User, error)
	Get(string) (*entity.User, error)

	Create(*entity.User) error
	Update(*entity.User, string) error
	Delete(string) error
}

// Delivery is
type Delivery interface {
	Fetch(echo.Context) error
	Get(echo.Context) error

	Create(echo.Context) error
	Update(echo.Context) error
	Delete(echo.Context) error
}
