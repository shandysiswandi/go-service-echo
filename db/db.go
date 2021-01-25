package db

import (
	"errors"
	"fmt"
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

	if config == nil {
		return db, append(errs, errors.New("Configuration is nil"))
	}

	if len(config.SchemaDatabases) < 1 {
		return db, append(errs, errors.New("This application not using any database"))
	}

	for _, s := range config.SchemaDatabases {
		if s == "mysql" && config.Gorm.MysqDSN != "" {
			db.Mysql, err = mysqlConnection(config.Gorm.MysqDSN)
			if err != nil {
				errs = append(errs, errors.New("Can't connect database mysql"))
			}
			continue
		}

		if s == "postgresql" && config.Gorm.PostgresqlDSN != "" {
			db.Posrgresql, err = postgresqlConnection(config.Gorm.PostgresqlDSN)
			if err != nil {
				errs = append(errs, errors.New("Can't connect database postgresql"))
			}
			continue
		}

		if s == "mongo" && config.Monggo.URI != "" && config.Monggo.Database != "" {
			db.Mongo, err = mongoConnection(config.Monggo.URI, config.Monggo.Database)
			if err != nil {
				errs = append(errs, errors.New("Can't connect database mongo"))
			}
			continue
		}
		errs = append(errs, fmt.Errorf("schema database '%s' is not support", s))
	}

	if len(errs) > 0 {
		return db, errs
	}

	return db, nil
}
