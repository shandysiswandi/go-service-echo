package entity

import "github.com/google/uuid"

// UserPayloadCreate is entity for validation on create
type UserPayloadCreate struct {
	Name     string `json:"name" validate:"required,min=5"`
	Email    string `json:"email" validate:"required,email,min=5"`
	Password string `json:"password" validate:"required,min=6"`
}

// User is
type User struct {
	UUID
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Task     []Task `json:"tasks,omitempty"`
	Timestamps
}

// TableName is
func (User) TableName() string {
	return "users"
}

// SetID is
func (u *User) SetID() {
	u.ID = uuid.New().String()
}
