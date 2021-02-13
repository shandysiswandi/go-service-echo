package config

import (
	"os"
	"strconv"
)

// New is
func New() *Config {
	config := new(Config)

	/* application configuration */
	config.App.Env = os.Getenv("APP_ENV")
	config.App.Port = os.Getenv("APP_PORT")
	config.App.Name = os.Getenv("APP_NAME")
	config.App.Timezone = os.Getenv("TZ")

	/* database configuration */
	config.Database.Driver = os.Getenv("DB_DRIVER")
	config.Database.Host = os.Getenv("DB_HOST")
	config.Database.Port = os.Getenv("DB_PORT")
	config.Database.Username = os.Getenv("DB_USERNAME")
	config.Database.Password = os.Getenv("DB_PASSWORD")
	config.Database.Name = os.Getenv("DB_NAME")

	/* external configuration */
	config.External.JsonplaceholderURL = os.Getenv("EXTERTNAL_JSONPLACEHOLDER_URL")

	/* library configuration */
	config.Library.SentryDSN = os.Getenv("LIBRARY_SENTRY_DSN")
	config.Library.JWT.AccessSecret = []byte(os.Getenv("LIBRARY_JWT_ACCESS_SECRET"))
	config.Library.JWT.RefreshSecret = []byte(os.Getenv("LIBRARY_JWT_REFRESH_SECRET"))
	config.Library.Redis.Addr = os.Getenv("LIBRARY_REDIS_ADDR")
	config.Library.Redis.Password = os.Getenv("LIBRARY_REDIS_PASSWORD")
	srd, err := strconv.Atoi(os.Getenv("LIBRARY_REDIS_DATABASE"))
	if err != nil {
		srd = 0
	}
	config.Library.Redis.Database = srd

	return config
}
