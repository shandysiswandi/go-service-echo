package task

import "github.com/labstack/echo/v4"

// NewRouter is
func NewRouter(e *echo.Echo) {
	d := NewDelivery()
	r := e.Group("/tasks")

	r.GET("", d.Fetch)
	r.GET("/:id", d.Get)
	r.POST("", d.Create)
	r.PUT("/:id", d.Update)
	r.DELETE("/:id", d.Delete)
}
