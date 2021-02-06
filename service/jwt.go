package service

import (
	"errors"
	"go-rest-echo/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type (
	// JWT is
	JWT struct {
		config *config.Config
	}

	// JWTToken is
	JWTToken struct {
		AccessToken  string
		RefreshToken string
	}

	// JWTClaimData is
	JWTClaimData struct {
		ID    string `json:"id"`
		Email string `json:"email"`
		Name  string `json:"name"`
	}

	// JWTClaim is
	JWTClaim struct {
		Data JWTClaimData `json:"data"`
		jwt.StandardClaims
	}
)

// NewJWT is
func NewJWT(c *config.Config) *JWT {
	return &JWT{config: c}
}

// Generate is
func (j *JWT) Generate(data JWTClaimData) (*JWTToken, error) {
	genToken := new(JWTToken)
	err := errors.New("Token Not Valid")

	// access token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &JWTClaim{
		data,
		jwt.StandardClaims{
			Audience:  "ACCESS_TOKEN_AUDIENCE",
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "ACCESS_TOKEN_ISSUER",
			Subject:   "ACCESS_TOKEN_SUBJECT",
		},
	})
	genToken.AccessToken, err = token.SignedString(j.config.Service.JWT.AccessSecret)
	if err != nil {
		return nil, err
	}

	// refresh token
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, &JWTClaim{
		data,
		jwt.StandardClaims{
			Audience:  "REFRESH_TOKEN_AUDIENCE",
			ExpiresAt: time.Now().Add(time.Minute * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "REFRESH_TOKEN_ISSUER",
			Subject:   "REFRESH_TOKEN_SUBJECT",
		},
	})
	genToken.RefreshToken, err = token.SignedString(j.config.Service.JWT.RefreshSecret)
	if err != nil {
		return nil, err
	}

	return genToken, nil
}
