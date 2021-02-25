package config_test

import (
	"go-service-echo/config"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew_App_Config(t *testing.T) {
	is := assert.New(t)

	os.Setenv("ENV", "testing")
	os.Setenv("HOST", "localhost")
	os.Setenv("PORT", "8080")
	os.Setenv("NAME", "go-service")
	os.Setenv("TZ", "UTC")

	actual := config.New().App

	is.Equal("testing", actual.Env)
	is.Equal("localhost", actual.Host)
	is.Equal("8080", actual.Port)
	is.Equal("go-service", actual.Name)
	is.Equal("UTC", actual.Timezone)
}

func TestNew_Database_Config(t *testing.T) {
	is := assert.New(t)

	os.Setenv("DB_DRIVER", "mysql")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_USERNAME", "root")
	os.Setenv("DB_PASSWORD", "password")
	os.Setenv("DB_NAME", "go-service")
	os.Setenv("DB_TIMEZONE", "UTC")

	actual := config.New().Database

	is.Equal("mysql", actual.Driver)
	is.Equal("localhost", actual.Host)
	is.Equal("3306", actual.Port)
	is.Equal("root", actual.Username)
	is.Equal("password", actual.Password)
	is.Equal("go-service", actual.Name)
	is.Equal("UTC", actual.Timezone)
}

func TestNew_JWT_Config(t *testing.T) {
	is := assert.New(t)

	os.Setenv("JWT_ACCESS_SECRET", "access")
	os.Setenv("JWT_REFRESH_SECRET", "refresh")

	actual := config.New().JWT

	is.Equal([]byte("access"), actual.AccessSecret)
	is.Equal([]byte("refresh"), actual.RefreshSecret)
}

func TestNew_Redis_Config(t *testing.T) {
	is := assert.New(t)

	os.Setenv("REDIS_ADDRESS", "localhost:1000")
	os.Setenv("REDIS_PASSWORD", "password")
	os.Setenv("REDIS_DATABASE", "0")

	actual := config.New().Redis

	is.Equal("localhost:1000", actual.Host)
	is.Equal("password", actual.Password)
	is.Equal(0, actual.Database)

	os.Setenv("REDIS_DATABASE", "a")
	actual = config.New().Redis
	is.Equal(0, actual.Database)
}

func TestNew_Sentry_Config(t *testing.T) {
	is := assert.New(t)

	os.Setenv("SENTRY_DSN", "dns")
	os.Setenv("SENTRY_ENV", "env")

	actual := config.New().Sentry

	is.Equal("dns", actual.DNS)
	is.Equal("env", actual.ENV)
}

func TestNew_External_Config(t *testing.T) {
	is := assert.New(t)

	os.Setenv("URL_JSONPLACEHOLDER", "url")

	actual := config.New().External

	is.Equal("url", actual.JSONPlaceHolder)
}
