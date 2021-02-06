package service

import (
	"go-rest-echo/config"
)

// Service is
type Service struct {
	Redis  Redis
	JWT    *JWT
	Logger *Logger
}

// New is
func New(c *config.Config) *Service {
	return &Service{
		Redis:  NewRedis(c),
		JWT:    NewJWT(c),
		Logger: NewLogger(c),
	}
}
