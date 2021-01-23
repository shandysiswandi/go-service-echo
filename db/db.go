package db

import (
	"errors"
	"go-rest-echo/config"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

// Database is
type Database struct {
	Mysql      *gorm.DB
	Posrgresql *gorm.DB
	Mongo      *mongo.Database
}

// NewDatabase is
func NewDatabase(config *config.Config) (*Database, []error) {
	var (
		err  error
		errs []error
		db   = new(Database)
	)

	// check SchemaDatabases is set or not
	if len(config.SchemaDatabases) < 1 {
		return nil, append(errs, errors.New("\033[33mThis application not using any database\033[0m"))
	}

	// loop SchemaDatabases and passing to database connection
	for _, s := range config.SchemaDatabases {
		if s == "mysql" {
			db.Mysql, err = mysqlConnection(config.Gorm.MysqDSN)
			if err != nil {
				errs = append(errs, err)
			}
			continue
		}

		if s == "postgresql" {
			db.Posrgresql, err = postgresqlConnection(config.Gorm.PostgresqlDSN)
			if err != nil {
				errs = append(errs, err)
			}
			continue
		}

		if s == "mongo" {
			db.Mongo, err = mongoConnection()
			if err != nil {
				errs = append(errs, err)
			}
			continue
		}
	}

	// check the errors
	if len(errs) > 0 {
		return db, errs
	}

	return db, nil
}
