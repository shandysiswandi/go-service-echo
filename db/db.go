package db

import (
	"errors"
	"fmt"
	"go-rest-echo/config"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

// all errros from db
var (
	ErrConfigIsNil          = errors.New("configuration is nil")
	ErrNotUseDatabase       = errors.New("this application not using any database")
	ErrNotConnectMysql      = errors.New("can't connect database mysql")
	ErrNotConnectPostgresql = errors.New("can't connect database postgresql")
	ErrNotConnectMongo      = errors.New("can't connect database mongo")
	ErrDriverNotSupport     = errors.New("driver database not support")
	ErrDialectNotSupport    = errors.New("dialect not support")
)

// all driver database support
var (
	MysqlDriver      = "mysql"
	PostgresqlDriver = "postgresql"
	MongoDriver      = "mongo"
)

// Database is
type Database struct {
	SQL   *gorm.DB
	Mongo *mongo.Database
}

// New is
func New(dc *config.DatabaseConfig, tz string) (*Database, error) {
	if dc == nil {
		return nil, ErrConfigIsNil
	}

	if dc.Driver == "" {
		return nil, ErrNotUseDatabase
	}

	var err error
	var db = &Database{}

	switch dc.Driver {
	case MysqlDriver:
		temp := "%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
		dsn := fmt.Sprintf(temp, dc.Username, dc.Password, dc.Host, dc.Port, dc.Name)

		db.SQL, err = gormConnection(dsn, MysqlDriver)
		if err != nil {
			return nil, ErrNotConnectMysql
		}
	case PostgresqlDriver:
		temp := "user=%s password=%s host=%s port=%s dbname=%s sslmode=disable TimeZone=%s"
		dsn := fmt.Sprintf(temp, dc.Username, dc.Password, dc.Host, dc.Port, dc.Name, tz)

		db.SQL, err = gormConnection(dsn, PostgresqlDriver)
		if err != nil {
			return nil, ErrNotConnectPostgresql
		}
	case MongoDriver:
		temp := "mongodb://%s:%s@%s:%s/?readPreference=primary&ssl=false"
		dsn := fmt.Sprintf(temp, dc.Username, dc.Password, dc.Host, dc.Port)

		db.Mongo, err = mongoConnection(dsn, dc.Name)
		if err != nil {
			return nil, ErrNotConnectMongo
		}
	default:
		return nil, ErrDriverNotSupport
	}

	return db, nil
}
