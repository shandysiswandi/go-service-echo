package blogs

import "github.com/google/uuid"

// BlogPayloadCreate is
type BlogPayloadCreate struct {
	Title string `json:"title" validate:"required,min=5"`
	Body  string `json:"body" validate:"required,min=15"`
}

// Blog is
type Blog struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

// TableName is
func (Blog) TableName() string {
	return "blogs"
}

// SetID is
func (u *Blog) SetID() {
	u.ID = uuid.New().String()
}
