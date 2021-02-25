package jwt

import (
	"errors"
	"go-service-echo/config"
	"time"

	lib "github.com/dgrijalva/jwt-go"
)

// all errors jwtlib
var (
	ErrGenerateAccessToken  = errors.New("failed generate access token")
	ErrGenerateRefreshToken = errors.New("failed generate refresh token")
)

// JWT is
type JWT struct {
	jwt *config.JWTConfig
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
	lib.StandardClaims
}

// New is
func New(c *config.JWTConfig) *JWT {
	return &JWT{c}
}

// GenerateAccessToken is
func (j *JWT) GenerateAccessToken(data ClaimData) (string, error) {
	token := lib.NewWithClaims(lib.SigningMethodHS256, &Claim{
		data, lib.StandardClaims{
			Audience:  "ACCESS_TOKEN_AUDIENCE",
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "ACCESS_TOKEN_ISSUER",
			Subject:   "ACCESS_TOKEN_SUBJECT",
		},
	})

	accessToken, err := token.SignedString(j.jwt.AccessSecret)
	if err != nil {
		return "", ErrGenerateAccessToken
	}

	return accessToken, nil
}

// GenerateRefreshToken is
func (j *JWT) GenerateRefreshToken(data ClaimData) (string, error) {
	token := lib.NewWithClaims(lib.SigningMethodHS256, &Claim{
		data, lib.StandardClaims{
			Audience:  "REFRESH_TOKEN_AUDIENCE",
			ExpiresAt: time.Now().Add(time.Minute * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "REFRESH_TOKEN_ISSUER",
			Subject:   "REFRESH_TOKEN_SUBJECT",
		},
	})

	refreshToken, err := token.SignedString(j.jwt.RefreshSecret)
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}
