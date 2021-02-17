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
type Pagination struct {
	Total    int `json:"total"`
	Limit    int `json:"limit"`
	Page     int `json:"page"`
	NextPage int `json:"next_page"`
	PrevPage int `json:"prev_page"`
}

// ResponseSuccessWithPaginate is
type ResponseSuccessWithPaginate struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
}
