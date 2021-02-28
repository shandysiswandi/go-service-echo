package context

import (
	"errors"
	"fmt"
	"go-service-echo/app/library/token"
	"go-service-echo/util/stringy"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// errors
var (
	ErrBadRequest          = errors.New("Bad Request, something wrong on your request")           // 400
	ErrInvalidCredential   = errors.New("Invalid Credential")                                     // 401
	ErrUnauthorized        = errors.New("Need Authorizetion Credential")                          // 401
	ErrNotFound            = errors.New("Not Found, your request data not found in our database") // 404
	ErrNotFoundRoute       = errors.New("Not Found, URL you want is not in this application")     // 404
	ErrMethodNotAllowed    = errors.New("Method Not Allowed")                                     // 405
	ErrUnprocessableEntity = errors.New("Validation Failed")                                      // 422
	ErrFailedGenerateToken = errors.New("Failed Generate Token")                                  // 500
	ErrInternalServer      = errors.New("Internal Server Error")                                  // 500

	// validation message
	minMsg      = "value must be at least"
	requiredMsg = "value must be required"
	emailMsg    = "value must be a valid email"
	defaultMsg  = "value must be validate"
)

type (
	// CustomContext is
	CustomContext struct {
		echo.Context
	}

	// ResponseError is
	ResponseError struct {
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Error   interface{} `json:"error"`
	}

	// ResponseSuccess is
	ResponseSuccess struct {
		Success bool        `json:"success"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}

	// Pagination is
	Pagination struct {
		Total    int `json:"total"`
		Limit    int `json:"limit"`
		Page     int `json:"page"`
		NextPage int `json:"next_page"`
		PrevPage int `json:"prev_page"`
	}

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
	e.HTTPErrorHandler = func(e error, c echo.Context) {
		code := http.StatusInternalServerError
		message := ErrInternalServer.Error()

		if he, ok := e.(*echo.HTTPError); ok {
			switch he.Code {
			case 404:
				message = ErrNotFoundRoute.Error()
				e = nil
			case 405:
				message = ErrMethodNotAllowed.Error()
				e = nil
			}

			code = he.Code
		}

		c.JSON(code, ResponseError{false, message, e})
	}
}

// ValidateVar is function to validate one line variable
// example:
// myEmail := "joeybloggs.gmail.com"
// if err := c.ValidateVar(id, "email"); err != nil {
//    return c.UnprocessableEntityVar(err)
// }
func (c *CustomContext) ValidateVar(value interface{}, tag string) map[string]interface{} {
	var v = validator.New()
	var e map[string]interface{}

	if err := v.Var(value, tag); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			e = map[string]interface{}{"message": fmt.Sprintf("`%v` %s", err.Value(), c.getMessageValidation(err))}
		}

		return e
	}

	return nil
}

// Success is | 200, 201, 204
func (c *CustomContext) Success(s int, m string, d interface{}) error {
	return c.JSON(s, ResponseSuccess{true, m, d})
}

// SuccessWithPaginate is | 200, 201, 204
func (c *CustomContext) SuccessWithPaginate(code int, m string, p Pagination, d interface{}) error {
	return c.JSON(code, ResponseSuccessWithPaginate{true, m, d, p})
}

// Unauthorized is | 401
func (c *CustomContext) Unauthorized(err interface{}) error {
	return c.commonError(http.StatusUnauthorized, ErrUnauthorized.Error(), err)
}

// BadRequest is | 400
func (c *CustomContext) BadRequest(err interface{}) error {
	return c.commonError(http.StatusBadRequest, ErrBadRequest.Error(), err)
}

// UnprocessableEntityVar is | 422
func (c *CustomContext) UnprocessableEntityVar(m map[string]interface{}) error {
	return c.commonError(http.StatusUnprocessableEntity, ErrUnprocessableEntity.Error(), m)
}

// UnprocessableEntity is | 422
func (c *CustomContext) UnprocessableEntity(err interface{}) error {
	var e []map[string]interface{}

	for _, err := range err.(validator.ValidationErrors) {
		e = append(e, map[string]interface{}{
			"key":     stringy.SnakeCase(err.StructField()),
			"message": c.getMessageValidation(err),
			"value":   err.Value(),
		})
	}

	return c.commonError(http.StatusUnprocessableEntity, ErrUnprocessableEntity.Error(), e)
}

// HandleErrors is | 4xx - 5xx
func (c *CustomContext) HandleErrors(e error) error {
	log.Println("Error:", e)

	// token expired | 400
	if errors.Is(e, token.ErrExpiredToken) {
		return c.commonError(http.StatusBadRequest, token.ErrExpiredToken.Error(), nil)
	}

	// token invalid 401
	if errors.Is(e, token.ErrInvalidToken) {
		return c.commonError(http.StatusUnauthorized, token.ErrInvalidToken.Error(), nil)
	}

	// no record in table of database | 404
	if errors.Is(e, gorm.ErrRecordNotFound) {
		return c.commonError(http.StatusNotFound, ErrNotFound.Error(), nil)
	}

	// login failed | 401
	if errors.Is(e, ErrInvalidCredential) {
		return c.commonError(http.StatusUnauthorized, ErrInvalidCredential.Error(), nil)
	}

	// login failed or failed generate token | 500
	if errors.Is(e, ErrFailedGenerateToken) {
		return c.commonError(http.StatusInternalServerError, ErrFailedGenerateToken.Error(), nil)
	}

	return c.commonError(http.StatusInternalServerError, ErrInternalServer.Error(), nil)
}

// commonError is
func (c *CustomContext) commonError(code int, m string, e interface{}) error {
	return c.JSON(code, ResponseError{false, m, e})
}

// private function to get message validation by flag | tag | reflect
func (c *CustomContext) getMessageValidation(e validator.FieldError) (msg string) {
	switch e.Tag() {
	case "min":
		msg = minMsg
	case "required":
		msg = requiredMsg
	case "email":
		msg = emailMsg
	default:
		msg = defaultMsg
	}

	if e.Param() != "" {
		msg = msg + " " + e.Param()
	}

	return msg
}
