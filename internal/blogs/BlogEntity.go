package blogs

import "github.com/google/uuid"

type (
	// BlogPayloadCreate is
	BlogPayloadCreate struct {
		Title string `json:"title" validate:"required,min=5"`
		Body  string `json:"body" validate:"required,min=15"`
	}

	// BlogPayloadPut is
	BlogPayloadPut struct {
		Title string `json:"title" validate:"required,min=5"`
		Body  string `json:"body" validate:"required,min=15"`
	}

	// BlogPayloadPatch is
	BlogPayloadPatch struct {
		Title string `json:"title,omitempty" validate:"omitempty,min=5"`
		Body  string `json:"body,omitempty" validate:"omitempty,min=15"`
	}

	// Blog is
	Blog struct {
		ID    string `json:"id"`
		Title string `json:"title"`
		Body  string `json:"body"`
	}
)

// TableName is
func (Blog) TableName() string {
	return "blogs"
}

// SetID is
func (u *Blog) SetID() {
	u.ID = uuid.New().String()
}
