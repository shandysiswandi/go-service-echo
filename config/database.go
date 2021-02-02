package config

import (
	"os"
	"strings"
)

type (
	mongo struct {
		URI string
		DB  string
	}

	// Database is
	database struct {
		Drivers       []string
		MysqlDSN      string
		PostgresqlDSN string
		Mongo         mongo
	}
)

func newDatabaseConfig() *database {
	return &database{
		Drivers:       strings.Split(os.Getenv("DB_DRIVERS"), ","),
		MysqlDSN:      os.Getenv("DB_MYSQL_DSN"),
		PostgresqlDSN: os.Getenv("DB_POSTGRESQL_DSN"),
		Mongo: mongo{
			URI: os.Getenv("DB_MONGO_URI"),
			DB:  os.Getenv("DB_MONGO_DATABASE"),
		},
	}
}
