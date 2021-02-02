package service

import (
	"go-rest-echo/config"
)

// Service is
type Service struct {
	Redis Redis
}

// New is
func New(c *config.Config) *Service {
	return &Service{
		Redis: NewRedis(c),
	}
}
