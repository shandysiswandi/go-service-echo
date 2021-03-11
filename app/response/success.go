package response

import "github.com/labstack/echo/v4"

// Success is
type Success struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// NewSuccess is
func NewSuccess(c echo.Context, code int, m string, d interface{}) error {
	return c.JSON(code, Success{
		Error:   false,
		Message: m,
		Data:    d,
	})
}
