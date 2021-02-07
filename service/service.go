package service

import (
	"go-rest-echo/config"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
)

type (
	// Service is
	Service struct {
		Redis  *Redis
		JWT    *JWT
		Sentry *Sentry
	}
	// JWT is
	JWT struct {
		config *config.Config
	}

	// JWTToken is
	JWTToken struct {
		AccessToken  string
		RefreshToken string
	}

	// JWTClaimData is
	JWTClaimData struct {
		ID    string `json:"id"`
		Email string `json:"email"`
		Name  string `json:"name"`
	}

	// JWTClaim is
	JWTClaim struct {
		Data JWTClaimData `json:"data"`
		jwt.StandardClaims
	}

	// Redis is service for caching
	Redis struct {
		client *redis.Client
	}

	// Sentry is
	Sentry struct {
		config *config.Config
	}
)

// New is
func New(c *config.Config) *Service {
	return &Service{
		Redis:  NewRedis(c),
		JWT:    NewJWT(c),
		Sentry: NewSentry(c),
	}
}
