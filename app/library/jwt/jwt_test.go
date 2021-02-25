package jwt_test

import (
	"go-service-echo/app/library/jwt"
	"go-service-echo/config"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJWT_Generate(t *testing.T) {
	is := assert.New(t)

	os.Setenv("JWT_ACCESS_SECRET", "secret")
	os.Setenv("JWT_REFRESH_SECRET", "refresh")

	theJWT := jwt.New(config.New().JWT)
	accessToken, err := theJWT.GenerateAccessToken(jwt.ClaimData{})
	refreshToken, err := theJWT.GenerateRefreshToken(jwt.ClaimData{})

	is.Nil(err)
	is.NotEqual("", accessToken)
	is.NotEqual("", refreshToken)
}
