package service

import (
	"go-rest-echo/config"

	"github.com/go-redis/redis"
)

type (
	// Redis is service for caching
	Redis interface {
		Get() error
		Set() error
	}

	redisClass struct {
		client *redis.Client
	}
)

// NewRedis is constructor
func NewRedis(c *config.Config) Redis {
	cl := redis.NewClient(&redis.Options{
		Addr:     c.Service.Redis.Addr,
		Password: c.Service.Redis.Password,
		DB:       c.Service.Redis.Database,
	})

	return &redisClass{
		client: cl,
	}
}

func (redisClass) Get() error {
	return nil
}

func (redisClass) Set() error {
	return nil
}
