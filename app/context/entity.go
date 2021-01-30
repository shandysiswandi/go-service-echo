package context

import "github.com/labstack/echo/v4"

type (
	// ResponseError is
	ResponseError struct {
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Error   interface{} `json:"error"`
	}

	// CustomContext is
	CustomContext struct {
		echo.Context
	}

	// ResponseSuccess is
	ResponseSuccess struct {
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	// Pagination is
	Pagination struct{}

	// ResponseSuccessWithPaginate is
	ResponseSuccessWithPaginate struct {
		Success    bool        `json:"success"`
		Message    string      `json:"message"`
		Data       interface{} `json:"data"`
		Pagination Pagination  `json:"pagination"`
	}
)
