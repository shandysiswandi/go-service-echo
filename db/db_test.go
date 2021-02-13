package db_test

import (
	"go-rest-echo/config"
	"go-rest-echo/db"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestNewDatabase_ConfigIsNil(t *testing.T) {
	is := assert.New(t)

	db, err := db.NewDatabase(nil)

	is.Nil(db)
	is.EqualError(err, "Configuration is nil")
}

func TestNewDatabase_DriverIsEmpty(t *testing.T) {
	is := assert.New(t)

	conf := &config.Config{}
	conf.Database.Driver = ""
	db, err := db.NewDatabase(conf)

	is.Nil(db)
	is.EqualError(err, "This application not using any database")
}

func TestNewDatabase_DriverMysqlButError(t *testing.T) {
	is := assert.New(t)

	conf := &config.Config{}
	conf.Database.Driver = "mysql"
	db, err := db.NewDatabase(conf)

	is.Nil(db)
	is.EqualError(err, "can't connect database mysql")
}

func TestNewDatabase_DriverPostgresqlButError(t *testing.T) {
	is := assert.New(t)

	conf := &config.Config{}
	conf.Database.Driver = "postgresql"
	db, err := db.NewDatabase(conf)

	is.Nil(db)
	is.EqualError(err, "can't connect database postgresql")
}

func TestNewDatabase_DriverMongoButError(t *testing.T) {
	is := assert.New(t)

	conf := &config.Config{}
	conf.Database.Driver = "mongo"
	db, err := db.NewDatabase(conf)

	is.Nil(db)
	is.EqualError(err, "can't connect database mongo")
}

func TestNewDatabase_DriverNotSupport(t *testing.T) {
	is := assert.New(t)

	conf := &config.Config{}
	conf.Database.Driver = "blabla"
	db, err := db.NewDatabase(conf)

	is.Nil(db)
	is.EqualError(err, "driver database not support")
}

func TestNewDatabase_MysqlConfigFromEnv(t *testing.T) {
	is := assert.New(t)

	if err := godotenv.Overload(".mysql"); err != nil {
		is.Nil(err)
	}

	db, err := db.NewDatabase(config.New())

	is.NotNil(db)
	is.NotNil(db.SQL)
	is.Nil(db.Mongo)
	is.Nil(err)
}

func TestNewDatabase_PostgresqlConfigFromEnv(t *testing.T) {
	is := assert.New(t)

	if err := godotenv.Overload(".postgresql"); err != nil {
		is.Nil(err)
	}

	db, err := db.NewDatabase(config.New())

	is.NotNil(db)
	is.NotNil(db.SQL)
	is.Nil(db.Mongo)
	is.Nil(err)
}

func TestNewDatabase_MongoConfigFromEnv(t *testing.T) {
	is := assert.New(t)

	if err := godotenv.Overload(".mongo"); err != nil {
		is.Nil(err)
	}

	log.Println(os.Getenv("DB_DRIVER"))

	db, err := db.NewDatabase(config.New())

	is.NotNil(db)
	is.NotNil(db.Mongo)
	is.Nil(db.SQL)
	is.Nil(err)
}
