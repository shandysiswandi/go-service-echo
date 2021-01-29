package service

import (
	"go-rest-echo/config"

	"github.com/go-redis/redis"
)

// Service is
type Service struct {
	Redis *redis.Client
}

// New is
func New(c *config.Config) *Service {
	return &Service{
		Redis: redisConnection(c),
	}
}

func redisConnection(c *config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     c.Service.Redis.Addr,
		Password: c.Service.Redis.Password,
		DB:       c.Service.Redis.Database,
	})
}
