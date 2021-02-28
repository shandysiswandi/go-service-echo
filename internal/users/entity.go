package users

import (
	"time"

	"github.com/google/uuid"
)

type (
	// UserRepository is
	UserRepository interface {
		Fetch() (Users, error)
		Get(string) (*User, error)
		GetByEmail(email string) (*User, error)

		Create(*User) error
		Update(*User, string) error
		Delete(string) error
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
func (u *User) SetID() {
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
