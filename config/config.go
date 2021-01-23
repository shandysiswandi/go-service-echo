package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// Config is
type Config struct {
	App struct {
		Env  string
		Port string
		Name string
	}
	SchemaDatabases []string
	Gorm            struct {
		MysqDSN       string
		PostgresqlDSN string
	}
	Monggo struct {
		URI      string
		Database string
	}
}

// NewConfiguration is
func NewConfiguration() (config *Config, err error) {
	err = godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	config.App.Env = os.Getenv("APP_ENV")
	config.App.Port = os.Getenv("APP_PORT")
	config.App.Name = os.Getenv("APP_NAME")

	config.SchemaDatabases = strings.Split(os.Getenv("SCHEMA_DATABASES"), ",")

	config.Gorm.MysqDSN = os.Getenv("GORM_MYSQL_DSN")
	config.Gorm.PostgresqlDSN = os.Getenv("GORM_POSTGRESQL_DSN")

	config.Monggo.URI = os.Getenv("MONGO_URI")
	config.Monggo.Database = os.Getenv("MONGO_DATABASE")

	return config, nil
}
