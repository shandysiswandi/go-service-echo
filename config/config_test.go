package config_test

import (
	"go-rest-echo/config"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestNewConfiguration(t *testing.T) {
	is := assert.New(t)

	if err := godotenv.Load(".env.test"); err != nil {
		is.Nil(err)
	}

	actual := config.New()

	is.Equal("testing", actual.App.Env)
	is.Equal("3000", actual.App.Port)
	is.Equal("go service", actual.App.Name)
	is.Equal("UTC", actual.App.Timezone)

	is.Equal("mysql", actual.Database.Driver)
	is.Equal("127.0.0.1", actual.Database.Host)
	is.Equal("3306", actual.Database.Port)
	is.Equal("root", actual.Database.Username)
	is.Equal("password", actual.Database.Password)
	is.Equal("go-service", actual.Database.Name)

	is.Equal("sentry_dsn", actual.Library.SentryDSN)
	is.Equal([]byte("access"), actual.Library.JWT.AccessSecret)
	is.Equal([]byte("refresh"), actual.Library.JWT.RefreshSecret)
	is.Equal("localhost:6379", actual.Library.Redis.Addr)
	is.Equal("your_password_redis", actual.Library.Redis.Password)
	is.Equal(0, actual.Library.Redis.Database)

	is.Equal("external_jsonplaceholder_url", actual.External.JsonplaceholderURL)
}
