package authentication

import (
	"go-rest-echo/app/context"
	"go-rest-echo/internal/users"
	"go-rest-echo/util"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type usecase struct {
	jwtScret       string
	userRepository users.UserRepository
}

// NewUsecase is
func NewUsecase(ur users.UserRepository, jwtScret string) AuthUsecase {
	return &usecase{userRepository: ur, jwtScret: jwtScret}
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

	claim := new(context.JwtClaims)
	claim.Email = user.Email
	claim.Name = user.Name
	claim.StandardClaims.ExpiresAt = time.Now().Add(time.Hour * 1000).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	t, err := token.SignedString([]byte(u.jwtScret))
	if err != nil {
		return nil, context.ErrFailedGenerateToken
	}

	rl := new(ResponseLogin)
	rl.Email = user.Email
	rl.Name = user.Name
	rl.Token = t

	return rl, nil
}
