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
	actual := config.NewConfiguration()

	assert.Equal(t, "env", actual.App.Env)
	assert.Equal(t, "port", actual.App.Port)
	assert.Equal(t, "name", actual.App.Name)

	assert.Equal(t, []string{"mysql", "postgresql", "mongo"}, actual.Database.Drivers)
	assert.Equal(t, "mysql_dsn", actual.Database.MysqlDSN)
	assert.Equal(t, "postgresql_dsn", actual.Database.PostgresqlDSN)
	assert.Equal(t, "mongo_uri", actual.Database.Mongo.URI)
	assert.Equal(t, "mongo_db", actual.Database.Mongo.DB)

	assert.Equal(t, "sentry_dsn", actual.Service.SentryDSN)
	assert.Equal(t, "addr", actual.Service.Redis.Addr)
	assert.Equal(t, "pass", actual.Service.Redis.Password)
	assert.Equal(t, 0, actual.Service.Redis.Database)

	assert.Equal(t, "external_jsonplaceholder_url", actual.External.JsonplaceholderURL)

	assert.Equal(t, "timezone", actual.Timezone)
	assert.Equal(t, "your_secret_for_jwt", actual.JwtSecret)
}
