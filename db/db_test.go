package db_test

import (
	"go-rest-echo/config"
	"go-rest-echo/db"
	"log"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestNewDatabase_ConfigIsNil(t *testing.T) {
	is := assert.New(t)

	db, errs := db.NewDatabase(nil)

	is.Nil(db)
	is.EqualError(errs[0], "Configuration is nil")
}

func TestNewDatabase_DriversIsEmpty(t *testing.T) {
	is := assert.New(t)

	conf := &config.Config{}
	conf.Database.Drivers = []string{}
	db, errs := db.NewDatabase(conf)

	is.Nil(db)
	is.EqualError(errs[0], "This application not using any database")
}

func TestNewDatabase_DriversNotEmptyButError(t *testing.T) {
	is := assert.New(t)

	conf := &config.Config{}
	conf.Database.Drivers = []string{"mysql", "postgresql", "mongo"}
	conf.Database.MysqlDSN = "username:password@tcp(127.0.0.1:3306)/database_name?charset=utf8mb4&parseTime=True&loc=Local"
	conf.Database.PostgresqlDSN = "user=username password=password dbname=database_name host=127.0.0.1 port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	conf.Database.Mongo.URI = "mongodb://username:password@127.0.0.1:27017/?readPreference=primary&ssl=false"
	conf.Database.Mongo.DB = "database_name"
	db, errs := db.NewDatabase(conf)

	is.NotNil(db)
	is.Nil(db.Mysql)
	is.Nil(db.Postgresql)
	is.Nil(db.Mongo)

	is.EqualError(errs[0], "Can't connect database mysql")
	is.EqualError(errs[1], "Can't connect database postgresql")
	is.EqualError(errs[2], "Can't connect database mongo")
}

func TestNewDatabase_DriversIsNotSupport(t *testing.T) {
	is := assert.New(t)

	conf := &config.Config{}
	conf.Database.Drivers = []string{"wa"}
	db, errs := db.NewDatabase(conf)

	is.NotNil(db)
	is.Nil(db.Mysql)
	is.Nil(db.Postgresql)
	is.Nil(db.Mongo)
	is.EqualError(errs[0], "schema database 'wa' is not support")
}

func TestNewDatabase(t *testing.T) {
	is := assert.New(t)

	if err := godotenv.Load(".env"); err != nil {
		log.Println(err)
	}

	actual, _ := db.NewDatabase(config.New())

	is.Nil(actual.Mysql.Error)
	is.Equal(int64(0), actual.Mysql.RowsAffected)

	is.Nil(nil, actual.Postgresql.Error)
	is.Equal(int64(0), actual.Postgresql.RowsAffected)

	is.Equal("go-rest-echo", actual.Mongo.Name())
}
