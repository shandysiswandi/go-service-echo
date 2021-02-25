package redis_test

import (
	"go-service-echo/app/library/redis"
	"go-service-echo/config"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	is := assert.New(t)

	os.Setenv("REDIS_HOST", "localhost:6790")
	os.Setenv("REDIS_PASSWORD", "")
	os.Setenv("REDIS_DATABASE", "0")

	// 1
	redis := redis.New(config.New().Redis)
	is.NotNil(redis)
	// 2
	err := redis.Get()
	is.Nil(err)
	// 3
	err = redis.Set()
	is.Nil(err)
	// 4
	err = redis.Increment()
	is.Nil(err)
	// 5
	err = redis.Decrement()
	is.Nil(err)
}
