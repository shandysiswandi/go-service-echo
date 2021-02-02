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

func TestNewConfiguration_Only_App(t *testing.T) {
	expected := new(config.Config)
	expected.App.Env = "env"
	expected.App.Port = "port"
	expected.App.Name = "name"

	got := config.NewConfiguration()
	assert.Equal(t, expected.App.Env, got.App.Env)
	assert.Equal(t, expected.App.Port, got.App.Port)
	assert.Equal(t, expected.App.Name, got.App.Name)
}

func TestNewConfiguration_Only_TZ(t *testing.T) {
	expected := new(config.Config)
	expected.Timezone = "timezone"

	got := config.NewConfiguration()
	assert.Equal(t, expected.Timezone, got.Timezone)
}

func TestNewConfiguration_Only_JWT(t *testing.T) {
	expected := new(config.Config)
	expected.JwtSecret = "your_secret_for_jwt"

	got := config.NewConfiguration()
	assert.Equal(t, expected.JwtSecret, got.JwtSecret)
}

func TestNewConfiguration_Only_Sentry(t *testing.T) {
	expected := new(config.Config)
	expected.SentryDSN = "sentry_dsn"

	got := config.NewConfiguration()
	assert.Equal(t, expected.SentryDSN, got.SentryDSN)
}

func TestNewConfiguration_Only_Service(t *testing.T) {
	expected := new(config.Config)
	expected.Service.Redis.Addr = "addr"
	expected.Service.Redis.Password = "pass"
	expected.Service.Redis.Database = 0

	actual := config.NewConfiguration()
	assert.Equal(t, expected.Service.Redis.Addr, actual.Service.Redis.Addr)
	assert.Equal(t, expected.Service.Redis.Password, actual.Service.Redis.Password)
	assert.Equal(t, expected.Service.Redis.Database, actual.Service.Redis.Database)
}

func TestNewConfiguration_Only_External(t *testing.T) {
	expected := new(config.Config)
	expected.External.JsonplaceholderURL = "external_jsonplaceholder_url"

	got := config.NewConfiguration()
	assert.Equal(t, expected.External.JsonplaceholderURL, got.External.JsonplaceholderURL)
}

func TestNewConfiguration_Only_Database(t *testing.T) {
	expected := new(config.Config)
	expected.Database.Drivers = []string{"mysql", "postgresql", "mongo"}
	expected.Database.MysqlDSN = "mysql_dsn"
	expected.Database.PostgresqlDSN = "postgresql_dsn"
	expected.Database.Mongo.URI = "mongo_uri"
	expected.Database.Mongo.DB = "mongo_db"

	got := config.NewConfiguration()
	assert.Equal(t, expected.Database.Drivers, got.Database.Drivers)
	assert.Equal(t, expected.Database.MysqlDSN, got.Database.MysqlDSN)
	assert.Equal(t, expected.Database.PostgresqlDSN, got.Database.PostgresqlDSN)
	assert.Equal(t, expected.Database.Mongo.URI, got.Database.Mongo.URI)
	assert.Equal(t, expected.Database.Mongo.DB, got.Database.Mongo.DB)
}
