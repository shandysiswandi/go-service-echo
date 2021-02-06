package db

import (
	"errors"
	"fmt"
	"go-rest-echo/config"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var (
	once sync.Once
	db   *Database
	err  error
	errs []error
)

// Database is
type Database struct {
	Mysql      *gorm.DB
	Postgresql *gorm.DB
	Mongo      *mongo.Database
}

// NewDatabase is
func NewDatabase(config *config.Config) (*Database, []error) {
	if config == nil {
		return nil, append(errs, errors.New("Configuration is nil"))
	}

	if len(config.Database.Drivers) < 1 {
		return nil, append(errs, errors.New("This application not using any database"))
	}

	once.Do(func() {
		db = new(Database)

		for _, s := range config.Database.Drivers {
			if s == "mysql" && config.Database.MysqlDSN != "" {
				db.Mysql, err = mysqlConnection(config.Database.MysqlDSN)
				if err != nil {
					errs = append(errs, errors.New("Can't connect database mysql"))
				}
				continue
			}

			if s == "postgresql" && config.Database.PostgresqlDSN != "" {
				db.Postgresql, err = postgresqlConnection(config.Database.PostgresqlDSN)
				if err != nil {
					errs = append(errs, errors.New("Can't connect database postgresql"))
				}
				continue
			}

			if s == "mongo" && config.Database.Mongo.URI != "" && config.Database.Mongo.DB != "" {
				db.Mongo, err = mongoConnection(config.Database.Mongo.URI, config.Database.Mongo.DB)
				if err != nil {
					errs = append(errs, errors.New("Can't connect database mongo"))
				}
				continue
			}

			errs = append(errs, fmt.Errorf("schema database '%s' is not support", s))
		}
	})

	if len(errs) > 0 {
		return db, errs
	}

	return db, nil
}
