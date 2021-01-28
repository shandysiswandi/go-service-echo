package config

import (
	"os"
	"strings"
)

// Config is
type Config struct {
	App struct {
		Env      string
		Port     string
		Name     string
		Timezone string
	}
	Database struct {
		Drivers       []string
		MysqlDSN      string
		PostgresqlDSN string
		Mongo         struct {
			URI string
			DB  string
		}
	}
	JwtSecret string
	SentryDSN string
	External  struct {
		JsonplaceholderURL string
	}
}

// NewConfiguration is
func NewConfiguration() *Config {
	config := new(Config)

	config.App.Env = os.Getenv("APP_ENV")
	config.App.Port = os.Getenv("APP_PORT")
	config.App.Name = os.Getenv("APP_NAME")
	config.App.Timezone = os.Getenv("APP_TZ")

	config.Database.Drivers = strings.Split(os.Getenv("DB_DRIVERS"), ",")
	config.Database.MysqlDSN = os.Getenv("DB_MYSQL_DSN")
	config.Database.PostgresqlDSN = os.Getenv("DB_POSTGRESQL_DSN")
	config.Database.Mongo.URI = os.Getenv("DB_MONGO_URI")
	config.Database.Mongo.DB = os.Getenv("DB_MONGO_DATABASE")

	config.JwtSecret = os.Getenv("JWT_SECRET")
	config.SentryDSN = os.Getenv("SENTRY_DSN")

	config.External.JsonplaceholderURL = os.Getenv("EXTERTNAL_JSONPLACEHOLDER_URL")

	return config
}
