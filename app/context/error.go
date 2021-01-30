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
)
