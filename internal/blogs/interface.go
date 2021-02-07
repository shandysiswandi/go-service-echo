package blogs

import "github.com/labstack/echo/v4"

type (
	// Repository is
	Repository interface {
		Fetch() (*[]Blog, error)
		Get(string) (*Blog, error)

		Create(*Blog) error
		Update(BlogPayloadPut, string) error
		UpdateField(BlogPayloadPatch, string) error
		Delete(string) error
	}

	// Usecase is
	Usecase interface {
		Fetch() (*[]Blog, error)
		Get(string) (*Blog, error)

		Create(BlogPayloadCreate) error
		Update(BlogPayloadPut, string) error
		UpdateField(BlogPayloadPatch, string) error
		Delete(string) error
	}

	// Delivery is
	Delivery interface {
		Fetch(echo.Context) error
		Get(echo.Context) error

		Create(echo.Context) error
		Update(echo.Context) error
		UpdateField(echo.Context) error
		Delete(echo.Context) error
	}
)
