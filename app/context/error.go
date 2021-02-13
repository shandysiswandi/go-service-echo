package context

import "errors"

// errors & errors message
var (
	ErrInvalidCredential   = errors.New("Invalid Credential")
	ErrFailedGenerateToken = errors.New("Failed Generate Token")

	ErrInvalidCredentialMessage   = "Invalid Credential"
	ErrFailedGenerateTokenMessage = "Failed Generate Token"
	ErrNotFoundMessage            = "Not Found, your request data not found in our database"
	ErrInternalServerMessage      = "Internal Server Error"
	ErrBadRequest                 = "Validation Failed"
	ErrUnprocessableEntity        = "Bad Request, something wrong on your request"

	// custom error
	err400 = "The URL you want is protected, you must supplied token"
	err401 = "The token you supplied is invalid"
	err404 = "The URL you want is not in this application"
	err405 = "The URL you want is not using this METHOD"

	// validation message
	minMsg      = "value must be at least"
	requiredMsg = "value must be required"
	emailMsg    = "value must be a valid email"
	defaultMsg  = "value must be validate"
)
