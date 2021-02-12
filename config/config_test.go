package config_test

import (
	"go-rest-echo/config"
	"log"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestNewConfiguration(t *testing.T) {
	if err := godotenv.Load(".env"); err != nil {
		log.Println(err)
	}

	is := assert.New(t)
	actual := config.NewConfiguration()

	is.Equal("testing", actual.App.Env)
	is.Equal("3000", actual.App.Port)
	is.Equal("go service", actual.App.Name)
	is.Equal("UTC", actual.App.Timezone)

	is.Equal([]string{"mysql"}, actual.Database.Drivers)
	is.Equal("mysql_dsn", actual.Database.MysqlDSN)
	is.Equal("postgresql_dsn", actual.Database.PostgresqlDSN)
	is.Equal("mongo_uri", actual.Database.Mongo.URI)
	is.Equal("mongo_db_name", actual.Database.Mongo.DB)

	is.Equal("sentry_dsn", actual.Service.SentryDSN)
	is.Equal([]byte("access"), actual.Service.JWT.AccessSecret)
	is.Equal([]byte("refresh"), actual.Service.JWT.RefreshSecret)
	is.Equal("localhost:6379", actual.Service.Redis.Addr)
	is.Equal("your_password_redis", actual.Service.Redis.Password)
	is.Equal(0, actual.Service.Redis.Database)

	is.Equal("external_jsonplaceholder_url", actual.External.JsonplaceholderURL)
}
