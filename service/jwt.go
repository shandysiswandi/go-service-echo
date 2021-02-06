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
)

// NewJWT is
func NewJWT(c *config.Config) *JWT {
	return &JWT{config: c}
}

// Generate is
func (j *JWT) Generate(data jwt.MapClaims) (*JWTToken, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	genToken := new(JWTToken)
	err := errors.New("Token Not Valid")

	// access token
	accClaim := token.Claims.(jwt.MapClaims)
	accClaim["data"] = data
	accClaim["aud"] = "ACCESS_TOKEN_AUDIENCE"
	accClaim["exp"] = time.Now().Add(time.Minute * 15).Unix()
	accClaim["iat"] = time.Now().Unix()
	accClaim["iss"] = "ACCESS_TOKEN"
	accClaim["sub"] = "ACCESS_TOKEN_SUBJECT"
	genToken.AccessToken, err = token.SignedString(j.config.Service.JWT.AccessSecret)
	if err != nil {
		return nil, err
	}

	// refresh token
	refClaim := token.Claims.(jwt.MapClaims)
	refClaim["data"] = data
	refClaim["aud"] = "REFRESH_TOKEN_AUDIENCE"
	refClaim["exp"] = time.Now().Add(time.Hour * 24).Unix()
	refClaim["iat"] = time.Now().Unix()
	refClaim["iss"] = "REFRESH_TOKEN"
	refClaim["sub"] = "REFRESH_TOKEN_SUBJECT"
	genToken.RefreshToken, err = token.SignedString(j.config.Service.JWT.RefreshSecret)
	if err != nil {
		return nil, err
	}

	return genToken, nil
}
