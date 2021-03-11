package response

import (
	"errors"
	"go-service-echo/app/library/token"
	"go-service-echo/util/stringy"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// errors response
var (
	ErrBadRequest          = errors.New("Bad Request, something wrong on your request")           // 400
	ErrInvalidCredential   = errors.New("Please provide valid credentials")                       // 401
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

// Error is
type Error struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Stack   interface{} `json:"stack"`
}

// NewError is
func NewError(c echo.Context, code int, m string, s interface{}) error {
	return c.JSON(code, Error{
		Error:   true,
		Message: m,
		Stack:   s,
	})
}

// BadRequest is | 400
func BadRequest(c echo.Context, s interface{}) error {
	return NewError(c, http.StatusBadRequest, ErrBadRequest.Error(), nil)
}

// Unauthorized is | 401
func Unauthorized(c echo.Context, s interface{}) error {
	return NewError(c, http.StatusUnauthorized, ErrUnauthorized.Error(), nil)
}

// UnprocessableEntityVar is
func UnprocessableEntityVar(c echo.Context, m map[string]interface{}) error {
	return NewError(c, http.StatusUnprocessableEntity, ErrUnprocessableEntity.Error(), m)
}

// UnprocessableEntity is
func UnprocessableEntity(c echo.Context, err interface{}) error {
	var e []map[string]interface{}

	for _, err := range err.(validator.ValidationErrors) {
		e = append(e, map[string]interface{}{
			"key":     stringy.SnakeCase(err.StructField()),
			"message": GetMessageValidation(err),
			"value":   err.Value(),
		})
	}

	return NewError(c, http.StatusUnprocessableEntity, ErrUnprocessableEntity.Error(), e)
}

// GetMessageValidation is
func GetMessageValidation(e validator.FieldError) (msg string) {
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

// HandleErrors is
func HandleErrors(c echo.Context, e error) error {
	println("Err: ", e)

	// token expired | 400
	if errors.Is(e, token.ErrExpiredToken) {
		return NewError(c, http.StatusBadRequest, token.ErrExpiredToken.Error(), nil)
	}

	// token invalid 401
	if errors.Is(e, token.ErrInvalidToken) {
		return NewError(c, http.StatusUnauthorized, token.ErrInvalidToken.Error(), nil)
	}

	// no record in table of database | 404
	if errors.Is(e, gorm.ErrRecordNotFound) {
		return NewError(c, http.StatusNotFound, ErrNotFound.Error(), nil)
	}

	// login failed | 401
	if errors.Is(e, ErrInvalidCredential) {
		return NewError(c, http.StatusUnauthorized, ErrInvalidCredential.Error(), nil)
	}

	// login failed or failed generate token | 500
	if errors.Is(e, ErrFailedGenerateToken) {
		return NewError(c, http.StatusInternalServerError, ErrFailedGenerateToken.Error(), nil)
	}

	// internal server error
	return NewError(c, http.StatusInternalServerError, ErrInternalServer.Error(), nil)
}

// DefaultEchoError is
func DefaultEchoError(e error, c echo.Context) {
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
	c.JSON(code, Error{true, message, e})
}
