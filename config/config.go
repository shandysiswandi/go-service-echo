package config

import (
	"os"
	"strconv"
)

// New is
func New() *Config {

	app := &AppConfig{
		Env:      os.Getenv("APP_ENV"),
		Port:     os.Getenv("APP_PORT"),
		Name:     os.Getenv("APP_NAME"),
		Timezone: os.Getenv("TZ"),
	}

	db := &DatabaseConfig{
		Driver:   os.Getenv("DB_DRIVER"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
	}

	jwt := &JWTConfig{
		AccessSecret:  []byte(os.Getenv("LIBRARY_JWT_ACCESS_SECRET")),
		RefreshSecret: []byte(os.Getenv("LIBRARY_JWT_REFRESH_SECRET")),
	}

	lrd, err := strconv.Atoi(os.Getenv("LIBRARY_REDIS_DATABASE"))
	if err != nil {
		lrd = 0
	}
	redis := &RedisConfig{
		Addr:     os.Getenv("LIBRARY_REDIS_ADDR"),
		Password: os.Getenv("LIBRARY_REDIS_PASSWORD"),
		Database: lrd,
	}

	lib := &LibraryConfig{
		SentryDSN: os.Getenv("LIBRARY_SENTRY_DSN"),
		JWT:       jwt,
		Redis:     redis,
	}

	ext := &ExternalConfig{
		JsonplaceholderURL: os.Getenv("EXTERTNAL_JSONPLACEHOLDER_URL"),
	}

	return &Config{app, db, lib, ext}
}
