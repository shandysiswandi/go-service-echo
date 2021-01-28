package db_test

import (
	"errors"
	"go-rest-echo/config"
	"go-rest-echo/db"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDatabase_ParamConfigIsNil(t *testing.T) {
	_, errs := db.NewDatabase(nil)

	assert.Equal(t, errors.New("Configuration is nil"), errs[0], "must be error")
}

func TestNewDatabase_ParamConfigDatabaseDriversIsEmpty(t *testing.T) {
	conf := new(config.Config)
	conf.Database.Drivers = []string{}

	_, errs := db.NewDatabase(conf)

	assert.Equal(t, errors.New("This application not using any database"), errs[0], "must be error")
}

func TestNewDatabase_ParamConfigDatabaseIsMysql_Error(t *testing.T) {
	conf := new(config.Config)
	conf.Database.Drivers = []string{"mysql"}
	conf.Database.MysqlDSN = "mysql"

	_, errs := db.NewDatabase(conf)

	assert.Equal(t, errors.New("Can't connect database mysql"), errs[0], "must be error")
}

func TestNewDatabase_ParamConfigDatabaseIsMysql_Success(t *testing.T) {
	conf := new(config.Config)
	conf.Database.Drivers = []string{"mysql"}
	conf.Database.MysqlDSN = "root:password@tcp(127.0.0.1:3500)/go-rest-echo?charset=utf8mb4&parseTime=True&loc=Local"

	want, _ := db.NewDatabase(nil)
	got, _ := db.NewDatabase(conf)

	assert.NotEqual(t, want.Mysql, got.Mysql, "must be success")
}

func TestNewDatabase_ParamConfigDatabaseIsPostgresql_Error(t *testing.T) {
	conf := new(config.Config)
	conf.Database.Drivers = []string{"postgresql"}
	conf.Database.PostgresqlDSN = "postgresql"

	_, errs := db.NewDatabase(conf)

	assert.Equal(t, errors.New("Can't connect database postgresql"), errs[0], "must be error")
}

func TestNewDatabase_ParamConfigDatabaseIsPostgresql_Success(t *testing.T) {
	conf := new(config.Config)
	conf.Database.Drivers = []string{"postgresql"}
	conf.Database.PostgresqlDSN = "user=root password=password dbname=go-rest-echo host=127.0.0.1 port=5300 sslmode=disable TimeZone=Asia/Jakarta"

	want, _ := db.NewDatabase(nil)
	got, _ := db.NewDatabase(conf)

	assert.NotEqual(t, want.Postgresql, got.Postgresql, "must be success")
}

func TestNewDatabase_ParamConfigDatabaseIsMongo_Error(t *testing.T) {
	conf := new(config.Config)
	conf.Database.Drivers = []string{"mongo"}
	conf.Database.Mongo.URI = "uri"
	conf.Database.Mongo.DB = "db"

	_, errs := db.NewDatabase(conf)

	assert.Equal(t, errors.New("Can't connect database mongo"), errs[0], "must be error")
}

func TestNewDatabase_ParamConfigDatabaseIsMongo_Success(t *testing.T) {
	conf := new(config.Config)
	conf.Database.Drivers = []string{"mongo"}
	conf.Database.Mongo.URI = "mongodb://root:password@localhost:27017/?readPreference=primary&ssl=false"
	conf.Database.Mongo.DB = "go-rest-echo"

	want, _ := db.NewDatabase(nil)
	got, _ := db.NewDatabase(conf)

	assert.NotEqual(t, want.Mongo, got.Mongo, "must be success")
}

func TestNewDatabase_ParamConfigDatabaseDriversAreNotSupport(t *testing.T) {
	conf := new(config.Config)
	conf.Database.Drivers = []string{"not-support-driver"}

	_, errs := db.NewDatabase(conf)

	assert.Equal(t, errors.New("schema database 'not-support-driver' is not support"), errs[0], "must be error")
}
