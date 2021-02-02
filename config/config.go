package config

import (
	"os"
	"sync"
)

var (
	once     sync.Once
	instance *Config
)

// Config is
type Config struct {
	App       *app
	Database  *database
	Service   *service
	External  *external
	Timezone  string
	JwtSecret string
}

// NewConfiguration is
func NewConfiguration() *Config {
	once.Do(func() {
		instance = new(Config)

		instance.App = newAppConfig()
		instance.Database = newDatabaseConfig()
		instance.Service = newServiceConfig()
		instance.External = newExternalConfig()

		instance.Timezone = os.Getenv("TZ")
		instance.JwtSecret = os.Getenv("JWT_SECRET")
	})

	return instance
}
