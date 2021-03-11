package users

import (
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type (
	// UserRepository is
	UserRepository interface {
		Fetch() (Users, error)
		Get(string) (*User, error)
		GetByEmail(string) (*User, error)

		Create(*UserCreatePayload) error
		Update(*UserUpdatePayload, string) error
		Delete(string) error
	}

	// UserUsecase is
	UserUsecase interface {
		Fetch() (Users, error)
		Get(string) (*User, error)
		GetByEmail(string) (*User, error)

		Create(*UserCreatePayload) error
		Update(*UserUpdatePayload, string) error
		Delete(string) error
	}

	// UserHandler is
	UserHandler interface {
		Fetch(echo.Context) error
		Get(echo.Context) error
		GetByEmail(echo.Context) error

		Create(echo.Context) error
		Update(echo.Context) error
		Delete(echo.Context) error
	}

	// UserResponse is
	UserResponse struct {
		ID        string     `json:"id"`
		Name      string     `json:"name"`
		Email     string     `json:"email"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt time.Time  `json:"updated_at"`
		DeletedAt *time.Time `json:"deleted_at"`
	}

	// UserCreatePayload is
	UserCreatePayload struct {
		ID        string    `json:"id,omitempty"`
		Name      string    `json:"name" validate:"required,min=5"`
		Email     string    `json:"email" validate:"required,email,min=5"`
		Password  string    `json:"password" validate:"required,min=6"`
		CreatedAt time.Time `json:"created_at,omitempty"`
		UpdatedAt time.Time `json:"updated_at,omitempty"`
	}

	// UserUpdatePayload is
	UserUpdatePayload struct {
		Name      string    `json:"name" validate:"required,min=5"`
		Email     string    `json:"email" validate:"required,email,min=5"`
		Password  string    `json:"password" validate:"required,min=6"`
		UpdatedAt time.Time `json:"updated_at,omitempty"`
	}

	// User is
	User struct {
		ID        string     `json:"id"`
		Name      string     `json:"name" validate:"required,min=5"`
		Email     string     `json:"email" validate:"required,email,min=5"`
		Password  string     `json:"password" validate:"required,min=6"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt time.Time  `json:"updated_at"`
		DeletedAt *time.Time `json:"deleted_at"`
	}

	// Users is
	Users []*User
)

// TableName is
func (User) TableName() string {
	return "users"
}

// SetID is
func (u *UserCreatePayload) SetID() {
	u.ID = uuid.New().String()
}

// Transform is
func (us *Users) Transform() []UserResponse {
	temp := []UserResponse{}
	for _, u := range *us {
		temp = append(temp, UserResponse{
			ID:        u.ID,
			Name:      u.Name,
			Email:     u.Email,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
			DeletedAt: u.DeletedAt,
		})
	}

	return temp
}
