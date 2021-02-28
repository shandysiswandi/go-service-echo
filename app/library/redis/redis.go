package redis

import (
	"go-service-echo/config"

	lib "github.com/go-redis/redis/v8"
)

// Redis is service for caching
type Redis struct {
	client *lib.Client
}

// New is constructor
func New(c *config.RedisConfig) *Redis {
	// dsn := fmt.Sprintf("redis://%s/<db>")
	// opt, err := lib.ParseURL("redis://localhost:6379/<db>")
	return &Redis{lib.NewClient(&lib.Options{
		Addr:     c.Host,
		Password: c.Password,
		DB:       c.Database,
	})}
}

// Get is
func (r *Redis) Get() error {
	// r.client.Get(1, "")
	return nil
}

// Set is
func (r *Redis) Set() error {
	return nil
}

// Increment is
func (r *Redis) Increment() error {
	return nil
}

// Decrement is
func (r *Redis) Decrement() error {
	return nil
}
