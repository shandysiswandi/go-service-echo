package context

import (
	"github.com/labstack/echo/v4"
)

// ResponseError is
type ResponseError struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}

// CustomContext is
type CustomContext struct {
	echo.Context
}

// ResponseSuccess is
type ResponseSuccess struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Pagination is
type Pagination struct{}

// ResponseSuccessWithPaginate is
type ResponseSuccessWithPaginate struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
}
