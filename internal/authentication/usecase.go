package authentication

import (
	"go-service-echo/app/context"
	"go-service-echo/app/library/token"
	"go-service-echo/internal/users"
	"go-service-echo/util/bcrypt"
	"time"
)

// AuthUsecase is
type AuthUsecase struct {
	token          *token.Token
	userRepository users.UserRepository
}

// NewUsecase is
func NewUsecase(ur users.UserRepository, token *token.Token) *AuthUsecase {
	return &AuthUsecase{token, ur}
}

// Login is
func (u *AuthUsecase) Login(pl *PayloadLogin) (*ResponseLogin, error) {
	user, err := u.userRepository.GetByEmail(pl.Email)
	if err != nil {
		return nil, context.ErrInvalidCredential
	}

	isPassValid := bcrypt.CheckPasswordHash(pl.Password, user.Password)
	if !isPassValid {
		return nil, context.ErrInvalidCredential
	}

	rl := new(ResponseLogin)
	rl.Email = user.Email
	rl.Name = user.Name

	// call generate token
	data := token.PayloadData{ID: user.ID, Email: user.Email}
	rl.AccessToken, err = u.token.NewAccessToken(data, time.Minute)
	if err != nil {
		return nil, err
	}

	rl.RefreshToken, err = u.token.NewRefreshToken(data, time.Hour)
	if err != nil {
		return nil, err
	}

	return rl, nil
}
