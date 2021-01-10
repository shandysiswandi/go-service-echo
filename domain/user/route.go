package user

import "github.com/labstack/echo/v4"

type router struct{}

// Router is
type Router interface {
	Initialize(*echo.Echo, string)
}

// NewRouter is
func NewRouter() Router {
	return &router{}
}

func (ro *router) Initialize(e *echo.Echo, p string) {
	r := e.Group(p)
	d := NewDelivery()
	r.GET("", d.Fetch)
	r.GET("/:id", d.Get)
	r.POST("", d.Create)
	r.PUT("/:id", d.Update)
	r.DELETE("/:id", d.Delete)
}
