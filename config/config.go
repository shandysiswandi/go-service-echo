package config

import (
	"os"
	"strconv"
)

// New is
func New() *Config {
	app := &AppConfig{
		Env:      os.Getenv("ENV"),
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		Name:     os.Getenv("NAME"),
		Timezone: os.Getenv("TZ"),
	}

	db := &DatabaseConfig{
		Driver:   os.Getenv("DB_DRIVER"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		Timezone: os.Getenv("DB_TIMEZONE"),
	}

	jwt := &JWTConfig{
		AccessSecret:  []byte(os.Getenv("JWT_ACCESS_SECRET")),
		RefreshSecret: []byte(os.Getenv("JWT_REFRESH_SECRET")),
	}

	lrd, err := strconv.Atoi(os.Getenv("REDIS_DATABASE"))
	if err != nil {
		lrd = 0
	}
	redis := &RedisConfig{
		Host:     os.Getenv("REDIS_ADDRESS"),
		Password: os.Getenv("REDIS_PASSWORD"),
		Database: lrd,
	}

	sentry := &SentryConfig{
		DNS: os.Getenv("SENTRY_DSN"),
		ENV: os.Getenv("SENTRY_ENV"),
	}

	ext := &ExternalConfig{
		JSONPlaceHolder: os.Getenv("URL_JSONPLACEHOLDER"),
	}

	return &Config{
		App:      app,
		Database: db,
		JWT:      jwt,
		Redis:    redis,
		Sentry:   sentry,
		External: ext,
	}
}
