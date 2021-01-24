package tasks

import (
	"github.com/labstack/echo/v4"
)

// TaskRepository is
type TaskRepository interface {
	Fetch() (*[]Task, error)
	Get(string) (*Task, error)

	Create(*Task) (*string, error)
	Update(*Task, string) error
	Delete(string) error
}

// TaskUsecase is
type TaskUsecase interface {
	Fetch() (*[]Task, error)
	Get(string) (*Task, error)

	Create(*Task) (*string, error)
	Update(*Task, string) error
	Delete(string) error
}

// TaskDelivery is
type TaskDelivery interface {
	Fetch(echo.Context) error
	Get(echo.Context) error

	Create(echo.Context) error
	Update(echo.Context) error
	Delete(echo.Context) error
}
