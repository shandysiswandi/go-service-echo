package context

import (
	"errors"
	"go-rest-echo/app/library/jwtlib"
	"go-rest-echo/util"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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

// GetJWT is
func (c *CustomContext) GetJWT() (*jwtlib.Claim, string) {
	user := c.Get("user").(*jwt.Token)
	return user.Claims.(*jwtlib.Claim), user.Raw
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

// HandleErrors is | 4xx - 5xx
func (c *CustomContext) HandleErrors(e error) error {
	log.Println("Error:", e)

	if errors.Is(e, gorm.ErrRecordNotFound) {
		return c.commonError(http.StatusNotFound, ErrNotFoundMessage, gorm.ErrRecordNotFound)
	}

	if errors.Is(e, ErrInvalidCredential) {
		return c.commonError(http.StatusUnauthorized, ErrInvalidCredentialMessage, ErrInvalidCredential)
	}

	if errors.Is(e, ErrFailedGenerateToken) {
		return c.commonError(http.StatusInternalServerError, ErrFailedGenerateTokenMessage, ErrFailedGenerateToken)
	}

	return c.commonError(http.StatusInternalServerError, ErrInternalServerMessage, e)
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
		msg = "value must be at least"
	case "required":
		msg = "value must be required"
	case "email":
		msg = "value must be a valid email"
	default:
		msg = "value must be validate"
	}

	if e.Param() != "" {
		msg = msg + " " + e.Param()
	}

	return msg
}
