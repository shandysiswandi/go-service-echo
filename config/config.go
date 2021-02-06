package config

import (
	"os"
	"strconv"
	"strings"
	"sync"
)

var (
	once     sync.Once
	instance *Config
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
	External struct {
		JsonplaceholderURL string
	}
	Service struct {
		JWT struct {
			AccessSecret  string
			RefreshSecret string
		}
		Redis struct {
			Addr     string
			Password string
			Database int
		}
		SentryDSN string
	}
}

// NewConfiguration is
func NewConfiguration() *Config {
	once.Do(func() {
		instance = new(Config)

		/* application configuration */
		instance.App.Env = os.Getenv("APP_ENV")
		instance.App.Port = os.Getenv("APP_PORT")
		instance.App.Name = os.Getenv("APP_NAME")
		instance.App.Timezone = os.Getenv("TZ")

		/* database configuration */
		instance.Database.Drivers = strings.Split(os.Getenv("DB_DRIVERS"), ",")
		instance.Database.MysqlDSN = os.Getenv("DB_MYSQL_DSN")
		instance.Database.PostgresqlDSN = os.Getenv("DB_POSTGRESQL_DSN")
		instance.Database.Mongo.URI = os.Getenv("DB_MONGO_URI")
		instance.Database.Mongo.DB = os.Getenv("DB_MONGO_DATABASE")

		/* external configuration */
		instance.External.JsonplaceholderURL = os.Getenv("EXTERTNAL_JSONPLACEHOLDER_URL")

		/* service configuration */
		// sentry for logging
		instance.Service.SentryDSN = os.Getenv("SERVICE_SENTRY_DSN")

		// jwt for authentification & authorization
		instance.Service.JWT.AccessSecret = os.Getenv("SERVICE_JWT_ACCESS_SECRET")
		instance.Service.JWT.RefreshSecret = os.Getenv("SERVICE_JWT_REFRESH_SECRET")

		// redis for caching
		instance.Service.Redis.Addr = os.Getenv("SERVICE_REDIS_ADDR")
		instance.Service.Redis.Password = os.Getenv("SERVICE_REDIS_PASSWORD")
		srd, err := strconv.Atoi(os.Getenv("SERVICE_REDIS_DATABASE"))
		if err != nil {
			srd = 0
		}
		instance.Service.Redis.Database = srd
	})

	return instance
}
