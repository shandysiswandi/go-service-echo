package redislib

import (
	"go-rest-echo/config"

	"github.com/go-redis/redis"
)

// Redis is service for caching
type Redis struct {
	client *redis.Client
}

// New is constructor
func New(c *config.Config) *Redis {
	return &Redis{redis.NewClient(&redis.Options{
		Addr:     c.Library.Redis.Addr,
		Password: c.Library.Redis.Password,
		DB:       c.Library.Redis.Database,
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

// Increment is
func (Redis) Increment() error {
	return nil
}

// Decrement is
func (Redis) Decrement() error {
	return nil
}
