package service

import (
	"go-rest-echo/config"

	"github.com/go-redis/redis"
)

type (
	// Redis is service for caching
	Redis struct {
		client *redis.Client
	}
)

// NewRedis is constructor
func NewRedis(c *config.Config) *Redis {
	return &Redis{redis.NewClient(&redis.Options{
		Addr:     c.Service.Redis.Addr,
		Password: c.Service.Redis.Password,
		DB:       c.Service.Redis.Database,
	})}
}

// Get is
func (Redis) Get() error {
	return nil
}

// Set is
func (Redis) Set() error {
	return nil
}
