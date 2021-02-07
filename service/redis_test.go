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

func TestRedis_Get_And_Set(t *testing.T) {
	c := config.NewConfiguration()
	redis := service.NewRedis(c)
	get := redis.Get()
	set := redis.Set()

	assert.NotEqual(t, nil, redis)
	assert.Equal(t, nil, get)
	assert.Equal(t, nil, set)
}
