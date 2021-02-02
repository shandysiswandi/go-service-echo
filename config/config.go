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
	Constant  *constant
	Database  *database
	External  *external
	Service   *service
	Timezone  string
	JwtSecret string
}

// NewConfiguration is
func NewConfiguration() *Config {
	once.Do(func() {
		instance = new(Config)

		instance.App = newAppConfig()
		instance.Constant = newConstantConfig()
		instance.Database = newDatabaseConfig()
		instance.External = newExternalConfig()
		instance.Service = newServiceConfig()

		instance.Timezone = os.Getenv("TZ")
		instance.JwtSecret = os.Getenv("JWT_SECRET")
	})

	return instance
}
