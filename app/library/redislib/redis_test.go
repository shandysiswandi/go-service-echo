package redislib_test

import (
	"go-rest-echo/app/library/redislib"
	"go-rest-echo/config"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	if err := godotenv.Load(".env"); err != nil {
		t.Error("no .env file")
	}

	is := assert.New(t)
	redis := redislib.New(config.New())

	is.Nil(redis)
}
