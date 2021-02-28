package config

import (
	"os"
	"strconv"
)

type (
	// Config is
	Config struct {
		App      *AppConfig
		SSL      *SSLConfig
		Database *DatabaseConfig
		Token    *TokenConfig
		Redis    *RedisConfig
		Sentry   *SentryConfig
		External *ExternalConfig
	}

	// AppConfig is
	AppConfig struct {
		Env      string
		Host     string
		Port     string
		Name     string
		Timezone string
	}

	// SSLConfig is
	SSLConfig struct {
		Cert string
		Key  string
	}

	// DatabaseConfig is
	DatabaseConfig struct {
		Driver   string
		Host     string
		Port     string
		Username string
		Password string
		Name     string
		Timezone string
	}

	// TokenConfig is
	TokenConfig struct {
		TokenType  string
		AccessKey  string
		RefreshKey string
	}

	// RedisConfig is
	RedisConfig struct {
		Host     string
		Password string
		Database int
	}

	// SentryConfig is
	SentryConfig struct {
		DNS string
		ENV string
	}

	// ExternalConfig is
	ExternalConfig struct {
		JSONPlaceHolder string
	}
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

	ssl := &SSLConfig{
		Cert: os.Getenv("SSL_CERT_PATH"),
		Key:  os.Getenv("SSL_KEY_PATH"),
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

	token := &TokenConfig{
		TokenType:  os.Getenv("TOKEN_TYPE"),
		AccessKey:  os.Getenv("TOKEN_ACCESS_KEY"),
		RefreshKey: os.Getenv("TOKEN_REFRESH_KEY"),
	}

	lrd, err := strconv.Atoi(os.Getenv("REDIS_DATABASE"))
	if err != nil {
		lrd = 0
	}
	redis := &RedisConfig{
		Host:     os.Getenv("REDIS_HOST"),
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
		SSL:      ssl,
		Database: db,
		Token:    token,
		Redis:    redis,
		Sentry:   sentry,
		External: ext,
	}
}
