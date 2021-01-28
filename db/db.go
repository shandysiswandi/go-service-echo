package db

import (
	"context"
	"errors"
	"fmt"
	"go-rest-echo/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database is
type Database struct {
	Mysql      *gorm.DB
	Postgresql *gorm.DB
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

	if len(config.Database.Drivers) < 1 {
		return db, append(errs, errors.New("This application not using any database"))
	}

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

	if len(errs) > 0 {
		return db, errs
	}

	return db, nil
}

// mysql connection
func mysqlConnection(dsn string) (db *gorm.DB, err error) {
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{PrepareStmt: true})
	if err != nil {
		return nil, err
	}

	// pooling connection
	sqlCon, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlCon.SetMaxIdleConns(10 / 2)
	sqlCon.SetMaxOpenConns(100 / 2)
	sqlCon.SetConnMaxLifetime(time.Hour / 2)

	return db, nil
}

// postgresql connection
func postgresqlConnection(dsn string) (db *gorm.DB, err error) {
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlCon, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlCon.SetMaxIdleConns(10 / 2)
	sqlCon.SetMaxOpenConns(100 / 2)
	sqlCon.SetConnMaxLifetime(time.Hour / 2)

	return db, nil
}

// mongo connection
func mongoConnection(uri, database string) (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client.Database(database), nil
}
