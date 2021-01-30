package context

import (
	"errors"
	"go-rest-echo/util"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

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

// New is constructor
func New(e *echo.Echo) {
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error { return next(&CustomContext{c}) }
	})

	// set custom error
	e.HTTPErrorHandler = httpErrorHandler
}

func httpErrorHandler(e error, c echo.Context) {
	code := http.StatusInternalServerError
	message := "Internal Server Error"

	if he, ok := e.(*echo.HTTPError); ok {
		switch he.Code {
		case 400:
			message = "The URL you want is protected, you must supplied token."
			e = errors.New("400")
			break
		case 401:
			message = "The token you supplied is invalid."
			e = errors.New("401")
			break
		case 404:
			message = "The URL you want is not in this application."
			e = errors.New("404")
			break
		case 405:
			message = "The URL you want is not using this METHOD."
			e = errors.New("405")
			break
		}

		code = he.Code
	}

	c.JSON(code, ResponseError{
		Success: false,
		Message: message,
		Error:   e,
	})
}

// Success is | 200, 201, 204
func (c *CustomContext) Success(s int, m string, d interface{}) error {
	return c.JSON(s, ResponseSuccess{
		Success: true,
		Message: m,
		Data:    d,
	})
}

// SuccessWithPaginate is | 200, 201, 204
func (c *CustomContext) SuccessWithPaginate(code int, m string, p Pagination, d interface{}) error {
	return c.JSON(code, ResponseSuccessWithPaginate{
		Success:    true,
		Message:    m,
		Pagination: p,
		Data:       d,
	})
}

// BadRequest is | 400
func (c *CustomContext) BadRequest(err interface{}) error {
	return c.commonError(
		http.StatusBadRequest,
		"Bad Request, something wrong on your request",
		err,
	)
}

// NotFound is | 404
func (c *CustomContext) NotFound(err interface{}) error {
	return c.commonError(
		http.StatusNotFound,
		"Not Found, your request data not found in our database",
		err,
	)
}

// UnprocessableEntity is | 422
func (c *CustomContext) UnprocessableEntity(err interface{}) error {
	var e []map[string]interface{}

	for _, err := range err.(validator.ValidationErrors) {
		e = append(e, map[string]interface{}{
			"key":     util.SnakeCase(err.StructField()),
			"message": c.getMessageValidation(err),
			"value":   err.Value(),
		})
	}

	return c.commonError(
		http.StatusUnprocessableEntity,
		"Validation Failed",
		e,
	)
}

// InternalServerError is | 500
func (c *CustomContext) InternalServerError(err interface{}) error {
	return c.commonError(
		http.StatusInternalServerError,
		"Internal Server Error",
		err,
	)
}

// HandleErrors is | 4xx - 50x
func (c *CustomContext) HandleErrors(e error) error {
	if errors.Is(e, gorm.ErrRecordNotFound) {
		return c.NotFound(gorm.ErrRecordNotFound)
	}

	if errors.Is(e, gorm.ErrInvalidTransaction) {
		return c.InternalServerError(gorm.ErrInvalidTransaction)
	}

	if errors.Is(e, gorm.ErrNotImplemented) {
		return c.InternalServerError(gorm.ErrNotImplemented)
	}

	if errors.Is(e, gorm.ErrMissingWhereClause) {
		return c.InternalServerError(gorm.ErrMissingWhereClause)
	}

	if errors.Is(e, gorm.ErrUnsupportedRelation) {
		return c.InternalServerError(gorm.ErrUnsupportedRelation)
	}

	if errors.Is(e, gorm.ErrPrimaryKeyRequired) {
		return c.InternalServerError(gorm.ErrPrimaryKeyRequired)
	}

	if errors.Is(e, gorm.ErrModelValueRequired) {
		return c.InternalServerError(gorm.ErrModelValueRequired)
	}

	if errors.Is(e, gorm.ErrInvalidData) {
		return c.InternalServerError(gorm.ErrInvalidData)
	}

	if errors.Is(e, gorm.ErrUnsupportedDriver) {
		return c.InternalServerError(gorm.ErrUnsupportedDriver)
	}

	if errors.Is(e, gorm.ErrRegistered) {
		return c.InternalServerError(gorm.ErrRegistered)
	}

	if errors.Is(e, gorm.ErrInvalidField) {
		return c.InternalServerError(gorm.ErrInvalidField)
	}

	if errors.Is(e, gorm.ErrEmptySlice) {
		return c.InternalServerError(gorm.ErrEmptySlice)
	}

	if errors.Is(e, gorm.ErrDryRunModeUnsupported) {
		return c.InternalServerError(gorm.ErrDryRunModeUnsupported)
	}

	return c.InternalServerError(e)
}

/*---------- private methods or functions ----------*/
func (c *CustomContext) commonError(code int, m string, e interface{}) error {
	return c.JSON(code, ResponseError{
		Success: false,
		Message: m,
		Error:   e,
	})
}

// private function to get message validation by flag | tag | reflect
func (c *CustomContext) getMessageValidation(e validator.FieldError) (msg string) {
	switch e.Tag() {
	case "min":
		msg = "value must be greater than"
	case "required":
		msg = "value must be required"
	default:
		msg = "value must be validate"
	}

	if e.Param() != "" {
		msg = msg + " " + e.Param()
	}

	return msg
}
