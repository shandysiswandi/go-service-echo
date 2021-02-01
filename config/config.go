package config

import (
	"os"
	"strconv"
	"strings"
	"sync"
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

var (
	once   sync.Once
	config *Config
)

// NewConfiguration is
func NewConfiguration() *Config {

	once.Do(func() {
		config = new(Config)

		config.App.Env = os.Getenv("APP_ENV")
		config.App.Port = os.Getenv("APP_PORT")
		config.App.Name = os.Getenv("APP_NAME")

		config.Timezone = os.Getenv("TZ")

		config.Database.Drivers = strings.Split(os.Getenv("DB_DRIVERS"), ",")
		config.Database.MysqlDSN = os.Getenv("DB_MYSQL_DSN")
		config.Database.PostgresqlDSN = os.Getenv("DB_POSTGRESQL_DSN")
		config.Database.Mongo.URI = os.Getenv("DB_MONGO_URI")
		config.Database.Mongo.DB = os.Getenv("DB_MONGO_DATABASE")

		config.JwtSecret = os.Getenv("JWT_SECRET")
		config.SentryDSN = os.Getenv("SENTRY_DSN")

		config.Service.Redis.Addr = os.Getenv("SERVICE_REDIS_ADDR")
		config.Service.Redis.Password = os.Getenv("SERVICE_REDIS_PASSWORD")
		config.Service.Redis.Database, _ = strconv.Atoi(os.Getenv("SERVICE_REDIS_DATABASE"))

		config.External.JsonplaceholderURL = os.Getenv("EXTERTNAL_JSONPLACEHOLDER_URL")
	})

	return config
}
