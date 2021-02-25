package redis

import (
	"go-service-echo/config"

	lib "github.com/go-redis/redis"
)

// Redis is service for caching
type Redis struct {
	client *lib.Client
}

// New is constructor
func New(c *config.RedisConfig) *Redis {
	return &Redis{lib.NewClient(&lib.Options{
		Addr:     c.Host,
		Password: c.Password,
		DB:       c.Database,
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
