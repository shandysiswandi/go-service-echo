package service_test

import (
	"go-rest-echo/config"
	"go-rest-echo/service"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func init() {
	err := godotenv.Load(".env.test")
	if err != nil {
		return
	}
}

func TestJWT_Generate(t *testing.T) {
	c := config.NewConfiguration()
	jwt := service.NewJWT(c)
	claim := service.JWTClaimData{}
	gen, err := jwt.Generate(claim)

	assert.Equal(t, nil, err)
	assert.NotNil(t, gen)
	assert.NotEqual(t, "", gen.AccessToken)
	assert.NotEqual(t, "", gen.RefreshToken)
}
