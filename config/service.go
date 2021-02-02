package config

import (
	"os"
	"strconv"
)

type (
	redis struct {
		Addr     string
		Password string
		Database int
	}

	// Service is
	service struct {
		Redis     redis
		SentryDSN string
	}
)

func newServiceConfig() *service {
	db, err := strconv.Atoi(os.Getenv("SERVICE_REDIS_DATABASE"))
	if err != nil {
		db = 0
	}

	return &service{
		Redis: redis{
			Addr:     os.Getenv("SERVICE_REDIS_ADDR"),
			Password: os.Getenv("SERVICE_REDIS_PASSWORD"),
			Database: db,
		},
		SentryDSN: os.Getenv("SENTRY_DSN"),
	}
}
