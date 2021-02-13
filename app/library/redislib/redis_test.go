package redislib_test

import (
	"go-rest-echo/app/library/redislib"
	"go-rest-echo/config"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	is := assert.New(t)

	if err := godotenv.Load(".redis"); err != nil {
		is.Nil(err)
	}

	redis := redislib.New(config.New())

	is.NotNil(redis)
}

func TestNew_Get(t *testing.T) {
	is := assert.New(t)

	if err := godotenv.Load(".redis"); err != nil {
		is.Nil(err)
	}

	redis := redislib.New(config.New())
	err := redis.Get()

	is.Nil(err)
}

func TestNew_Set(t *testing.T) {
	is := assert.New(t)

	if err := godotenv.Load(".redis"); err != nil {
		is.Nil(err)
	}

	redis := redislib.New(config.New())
	err := redis.Set()

	is.Nil(err)
}

func TestNew_Increment(t *testing.T) {
	is := assert.New(t)

	if err := godotenv.Load(".redis"); err != nil {
		is.Nil(err)
	}

	redis := redislib.New(config.New())
	err := redis.Increment()

	is.Nil(err)
}

func TestNew_Decrement(t *testing.T) {
	is := assert.New(t)

	if err := godotenv.Load(".redis"); err != nil {
		is.Nil(err)
	}

	redis := redislib.New(config.New())
	err := redis.Decrement()

	is.Nil(err)
}
