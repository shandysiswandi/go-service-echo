package authentication

type (
	// PayloadLogin is
	PayloadLogin struct {
		Email    string `json:"email" validate:"required,email,min=5"`
		Password string `json:"password" validate:"required,min=8"`
	}

	// ResponseLogin is
	ResponseLogin struct {
		Email        string `json:"email"`
		Name         string `json:"name"`
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}
)
