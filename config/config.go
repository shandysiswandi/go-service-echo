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
		Env  string
		Port string
		Name string
	}
	Timezone string
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
	Service   struct {
		Redis struct {
			Addr     string
			Password string
			Database int
		}
	}
	External struct {
		JsonplaceholderURL string
	}
}

// NewConfiguration is
func NewConfiguration() *Config {
	once.Do(func() {
		instance = new(Config)

		instance.App.Env = os.Getenv("APP_ENV")
		instance.App.Port = os.Getenv("APP_PORT")
		instance.App.Name = os.Getenv("APP_NAME")

		instance.Timezone = os.Getenv("TZ")

		instance.Database.Drivers = strings.Split(os.Getenv("DB_DRIVERS"), ",")
		instance.Database.MysqlDSN = os.Getenv("DB_MYSQL_DSN")
		instance.Database.PostgresqlDSN = os.Getenv("DB_POSTGRESQL_DSN")
		instance.Database.Mongo.URI = os.Getenv("DB_MONGO_URI")
		instance.Database.Mongo.DB = os.Getenv("DB_MONGO_DATABASE")

		instance.JwtSecret = os.Getenv("JWT_SECRET")
		instance.SentryDSN = os.Getenv("SENTRY_DSN")

		instance.Service.Redis.Addr = os.Getenv("SERVICE_REDIS_ADDR")
		instance.Service.Redis.Password = os.Getenv("SERVICE_REDIS_PASSWORD")
		instance.Service.Redis.Database, _ = strconv.Atoi(os.Getenv("SERVICE_REDIS_DATABASE"))

		instance.External.JsonplaceholderURL = os.Getenv("EXTERTNAL_JSONPLACEHOLDER_URL")
	})

	return instance
}
