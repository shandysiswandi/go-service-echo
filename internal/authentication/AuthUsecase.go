package authentication

import (
	"go-rest-echo/app/context"
	"go-rest-echo/app/library/jwtlib"
	"go-rest-echo/internal/users"
	"go-rest-echo/util/bcrypt"
)

type usecase struct {
	jwt            *jwtlib.JWT
	userRepository users.UserRepository
}

// NewUsecase is
func NewUsecase(ur users.UserRepository, jwt *jwtlib.JWT) AuthUsecase {
	return &usecase{jwt, ur}
}

func (u *usecase) Login(pl *PayloadLogin) (*ResponseLogin, error) {
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
	data := jwtlib.ClaimData{ID: user.ID, Email: user.Email, Name: user.Name}
	rl.AccessToken, rl.RefreshToken, err = u.jwt.Generate(data)
	if err != nil {
		return nil, err
	}

	return rl, nil
}
