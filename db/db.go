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
	SQL   *gorm.DB
	Mongo *mongo.Database
}

// NewDatabase is
func NewDatabase(c *config.Config) (*Database, error) {
	if c == nil {
		return nil, errors.New("Configuration is nil")
	}

	if c.Database.Driver == "" {
		return nil, errors.New("This application not using any database")
	}

	var err error
	var db = &Database{}

	switch c.Database.Driver {
	case "mysql":
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			c.Database.Username,
			c.Database.Password,
			c.Database.Host,
			c.Database.Port,
			c.Database.Name,
		)
		db.SQL, err = mysqlConnection(dsn)
		if err != nil {
			return nil, errors.New("can't connect database mysql")
		}
	case "postgresql":
		dsn := fmt.Sprintf(
			"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable TimeZone=%s",
			c.Database.Username,
			c.Database.Password,
			c.Database.Host,
			c.Database.Port,
			c.Database.Name,
			c.App.Timezone,
		)
		db.SQL, err = postgresqlConnection(dsn)
		if err != nil {
			return nil, errors.New("can't connect database postgresql")
		}
	case "mongo":
		uri := fmt.Sprintf(
			"mongodb://%s:%s@%s:%s/?readPreference=primary&ssl=false",
			c.Database.Username,
			c.Database.Password,
			c.Database.Host,
			c.Database.Port,
		)
		db.Mongo, err = mongoConnection(uri, c.Database.Name)
		if err != nil {
			return nil, errors.New("can't connect database mongo")
		}
	default:
		return nil, errors.New("driver database not support")
	}

	return db, nil
}
