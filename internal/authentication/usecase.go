package authentication

import (
	"go-rest-echo/app/context"
	"go-rest-echo/internal/users"
	"go-rest-echo/service"
	"go-rest-echo/util"
)

type usecase struct {
	service        *service.Service
	userRepository users.UserRepository
}

// NewUsecase is
func NewUsecase(ur users.UserRepository, service *service.Service) AuthUsecase {
	return &usecase{userRepository: ur, service: service}
}

func (u *usecase) Login(pl *PayloadLogin) (*ResponseLogin, error) {
	user, err := u.userRepository.GetByEmail(pl.Email)
	if err != nil {
		return nil, context.ErrInvalidCredential
	}

	isPassValid := util.CheckPasswordHash(pl.Password, user.Password)
	if !isPassValid {
		return nil, context.ErrInvalidCredential
	}

	// call generate token
	tok, err := u.service.JWT.Generate(service.JWTClaimData{ID: user.ID, Email: user.Email, Name: user.Name})
	if err != nil {
		return nil, err
	}

	rl := new(ResponseLogin)
	rl.Email = user.Email
	rl.Name = user.Name
	rl.AccessToken = tok.AccessToken
	rl.RefreshToken = tok.RefreshToken

	return rl, nil
}
