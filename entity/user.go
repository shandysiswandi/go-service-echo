package entity

// User is
type User struct {
	UUID
	Name     string `json:"name" validate:"required,min=5"`
	Email    string `json:"email" validate:"required,email,min=5"`
	Password string `json:"password" validate:"required,min=6"`
	Task     []Task `json:"tasks"`
	Timestamps
}

// TableName is
func (User) TableName() string {
	return "users"
}
