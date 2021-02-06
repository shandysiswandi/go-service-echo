package config_test

import (
	"go-rest-echo/config"
	"log"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func init() {
	err := godotenv.Load(".env.test")
	if err != nil {
		log.Println(err)
		return
	}
}

func TestNewConfiguration(t *testing.T) {
	ass := assert.New(t)
	actual := config.NewConfiguration()

	ass.Equal("env", actual.App.Env)
	ass.Equal("port", actual.App.Port)
	ass.Equal("name", actual.App.Name)
	ass.Equal("timezone", actual.App.Timezone)

	ass.Equal([]string{"mysql", "postgresql", "mongo"}, actual.Database.Drivers)
	ass.Equal("mysql_dsn", actual.Database.MysqlDSN)
	ass.Equal("postgresql_dsn", actual.Database.PostgresqlDSN)
	ass.Equal("mongo_uri", actual.Database.Mongo.URI)
	ass.Equal("mongo_db", actual.Database.Mongo.DB)

	ass.Equal("sentry_dsn", actual.Service.SentryDSN)
	ass.Equal([]byte("access"), actual.Service.JWT.AccessSecret)
	ass.Equal([]byte("refresh"), actual.Service.JWT.RefreshSecret)
	ass.Equal("addr", actual.Service.Redis.Addr)
	ass.Equal("pass", actual.Service.Redis.Password)
	ass.Equal(0, actual.Service.Redis.Database)

	ass.Equal("external_jsonplaceholder_url", actual.External.JsonplaceholderURL)
}
