package helper

import "github.com/google/uuid"

// GenerateUUID is
func GenerateUUID() string {
	return uuid.New().String()
}
