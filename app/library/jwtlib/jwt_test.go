package jwtlib_test

import (
	"go-rest-echo/app/library/jwtlib"
	"go-rest-echo/config"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestJWT_Generate(t *testing.T) {
	if err := godotenv.Load(".env"); err != nil {
		t.Error("no .env file")
	}

	is := assert.New(t)
	jwt := jwtlib.New(config.New())
	accessToken, refreshToken, err := jwt.Generate(jwtlib.ClaimData{})

	is.Nil(err)
	is.NotEqual("", accessToken)
	is.NotEqual("", refreshToken)
}
