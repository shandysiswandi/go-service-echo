package config

import (
	"os"
	"strconv"
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
	External struct {
		JsonplaceholderURL string
	}
	Service struct {
		JWT struct {
			AccessSecret  []byte
			RefreshSecret []byte
		}
		Redis struct {
			Addr     string
			Password string
			Database int
		}
		SentryDSN string
	}
}

// New is
func New() *Config {
	config := new(Config)

	/* application configuration */
	config.App.Env = os.Getenv("APP_ENV")
	config.App.Port = os.Getenv("APP_PORT")
	config.App.Name = os.Getenv("APP_NAME")
	config.App.Timezone = os.Getenv("TZ")

	/* database configuration */
	config.Database.Drivers = getDBDriver()
	config.Database.MysqlDSN = os.Getenv("DB_MYSQL_DSN")
	config.Database.PostgresqlDSN = os.Getenv("DB_POSTGRESQL_DSN")
	config.Database.Mongo.URI = os.Getenv("DB_MONGO_URI")
	config.Database.Mongo.DB = os.Getenv("DB_MONGO_DATABASE")

	/* external configuration */
	config.External.JsonplaceholderURL = os.Getenv("EXTERTNAL_JSONPLACEHOLDER_URL")

	/* service configuration */
	// sentry for logging online
	config.Service.SentryDSN = os.Getenv("SERVICE_SENTRY_DSN")

	// jwt for authentification & authorization
	config.Service.JWT.AccessSecret = []byte(os.Getenv("SERVICE_JWT_ACCESS_SECRET"))
	config.Service.JWT.RefreshSecret = []byte(os.Getenv("SERVICE_JWT_REFRESH_SECRET"))

	// redis for caching
	config.Service.Redis.Addr = os.Getenv("SERVICE_REDIS_ADDR")
	config.Service.Redis.Password = os.Getenv("SERVICE_REDIS_PASSWORD")
	config.Service.Redis.Database = getServiceRedisDatabase()

	return config
}

func getDBDriver() []string {
	var drivers []string
	split := strings.Split(os.Getenv("DB_DRIVERS"), ",")

	for _, s := range split {
		if strings.TrimSpace(s) != "" {
			drivers = append(drivers, s)
		}
	}

	return drivers
}

func getServiceRedisDatabase() int {
	srd, err := strconv.Atoi(os.Getenv("SERVICE_REDIS_DATABASE"))
	if err != nil {
		srd = 0
	}

	return srd
}
