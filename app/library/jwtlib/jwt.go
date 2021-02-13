package jwtlib

import (
	"errors"
	"go-rest-echo/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWT is
type JWT struct {
	config *config.Config
}

// ClaimData is
type ClaimData struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

// Claim is
type Claim struct {
	Data ClaimData `json:"data"`
	jwt.StandardClaims
}

// New is
func New(c *config.Config) *JWT {
	return &JWT{config: c}
}

// Generate is
func (j *JWT) Generate(data ClaimData) (string, string, error) {
	err := errors.New("token not valid")

	accessToken, err := j.accessToken(data, j.config.Library.JWT.AccessSecret)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := j.refreshToken(data, j.config.Library.JWT.RefreshSecret)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (JWT) accessToken(data ClaimData, secret []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claim{
		data, jwt.StandardClaims{
			Audience:  "ACCESS_TOKEN_AUDIENCE",
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "ACCESS_TOKEN_ISSUER",
			Subject:   "ACCESS_TOKEN_SUBJECT",
		},
	})

	accessToken, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (JWT) refreshToken(data ClaimData, secret []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claim{
		data, jwt.StandardClaims{
			Audience:  "REFRESH_TOKEN_AUDIENCE",
			ExpiresAt: time.Now().Add(time.Minute * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "REFRESH_TOKEN_ISSUER",
			Subject:   "REFRESH_TOKEN_SUBJECT",
		},
	})

	refreshToken, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}
