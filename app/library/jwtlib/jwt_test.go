package jwtlib_test

import (
	"go-rest-echo/app/library/jwtlib"
	"go-rest-echo/config"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestJWT_Generate(t *testing.T) {
	is := assert.New(t)

	if err := godotenv.Load(".jwt"); err != nil {
		is.Nil(err)
	}

	jwt := jwtlib.New(config.New().Library.JWT)
	accessToken, err := jwt.GenerateAccessToken(jwtlib.ClaimData{})
	refreshToken, err := jwt.GenerateRefreshToken(jwtlib.ClaimData{})

	is.Nil(err)
	is.NotEqual("", accessToken)
	is.NotEqual("", refreshToken)
}
