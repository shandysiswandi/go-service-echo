package service

import (
	"go-rest-echo/config"
)

// Service is
type Service struct {
	Redis  Redis
	JWT    *JWT
	Sentry *Sentry
	Logger *Logger
}

// New is
func New(c *config.Config) *Service {
	return &Service{
		Redis:  NewRedis(c),
		JWT:    NewJWT(c),
		Sentry: NewSentry(c),
		Logger: NewLogger(c),
	}
}
