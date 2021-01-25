package db_test

import (
	"errors"
	"fmt"
	"go-rest-echo/config"
	"go-rest-echo/db"
	"testing"

	"github.com/stretchr/testify/assert"
)

/* Skenario tes fungsi NewDatabase adalah sebagai berikut:
 *
 * 1. cek param `*config.Config`, apakah bernilai `nil`
 * 	a. ya, maka akan mengembalikan `new(db.Database)` dan `[]error index [0] = errors.New("Configuration is nil")`
 *	b. tidak, maka lanjutkan
 *
 * 2. cek panjang *config.Config.SchemaDatabases, apakah kurang dari satu
 * 	a. ya, maka akan mengembalikan `new(Database)` dan `[]error index [0] = errors.New("This application not using any database")`
 *	b. tidak, maka lanjutkan
 *
 * 3. looping *config.Config.SchemaDatabases, lalu cek dan bandingkan apakah value sama dengan antara `mysql`,`postgresql`,`mongo`
 * 	a. jika sama dengan `mysql`, maka jika tidak ada error
 *		return from connection = `*db.Database.Mysql`, `nil` lalu continue, jika ada error maka append ke `errs`
 *
 * 	b. jika sama dengan `postgresql`, maka jika tidak ada error
 *		return from connection = `*db.Database.Postgresql`, `nil` lalu continue, jika ada error maka append ke `errs`
 *
 * 	c. jika sama dengan `mongo`, maka jika tidak ada error
 *		return from connection = `*db.Database.Mongo`, `nil` lalu continue, jika ada error maka append ke `errs`
 *
 *	d. jika tidak diantara `mysql`,`postgresql`,`mongo`, maka akan return `db`, `errs`
 *
 *	e. setelah looping selesai, maka check panjang `errs`, jika lebih dari 0 maka return `db`, `errs`, jika tidak lanjutkan
 *
 * 4. return `db` and `nil`
 */

func TestNewDatabase_ParamConfigIsNil(t *testing.T) {
	want := new(db.Database)
	got, errs := db.NewDatabase(nil)

	assert.Equal(t, []error{errors.New("Configuration is nil")}, errs)
	assert.Equal(t, errors.New("Configuration is nil"), errs[0])
	assert.Equal(t, want, got)
}

func TestNewDatabase_ParamConfigSchemaDatabasesIsEmpty(t *testing.T) {
	conf := &config.Config{
		App: struct {
			Env  string
			Port string
			Name string
		}{
			"env",
			"port",
			"name",
		},
		SchemaDatabases: []string{},
		Gorm: struct {
			MysqDSN       string
			PostgresqlDSN string
		}{
			"MysqDSN",
			"PostgresqlDSN",
		},
		Monggo: struct {
			URI      string
			Database string
		}{
			"URI",
			"Database",
		},
	}
	want := new(db.Database)
	got, errs := db.NewDatabase(conf)

	assert.Equal(t, []error{errors.New("This application not using any database")}, errs)
	assert.Equal(t, errors.New("This application not using any database"), errs[0])
	assert.Equal(t, want, got)
}

func TestNewDatabase_ParamConfigSchemaDatabasesIsMysql(t *testing.T) {
	conf := &config.Config{
		App: struct {
			Env  string
			Port string
			Name string
		}{
			"env",
			"port",
			"name",
		},
		SchemaDatabases: []string{"mysql"},
		Gorm: struct {
			MysqDSN       string
			PostgresqlDSN string
		}{
			"root:password@tcp(127.0.0.1:3500)/go-rest-echo111?charset=utf8mb4&parseTime=True&loc=Local",
			"user=root password=password dbname=go-rest-echo host=127.0.0.1 port=5300 sslmode=disable TimeZone=Asia/Jakarta",
		},
		Monggo: struct {
			URI      string
			Database string
		}{
			"mongodb://root:password@localhost:27017/?readPreference=primary&ssl=false",
			"go-rest-echo",
		},
	}

	// test if error on connection
	want := new(db.Database)
	got, errs := db.NewDatabase(conf)

	assert.Equal(t, []error{errors.New("Can't connect database mysql")}, errs)
	assert.Equal(t, want, got)

	// test if not error
	conf.Gorm.MysqDSN = "root:password@tcp(127.0.0.1:3500)/go-rest-echo?charset=utf8mb4&parseTime=True&loc=Local"
	got, _ = db.NewDatabase(conf)

	assert.NotEqual(t, nil, got.Mysql)
}

func TestNewDatabase_ParamConfigSchemaDatabasesIsPostgresql(t *testing.T) {
	conf := &config.Config{
		App: struct {
			Env  string
			Port string
			Name string
		}{
			"env",
			"port",
			"name",
		},
		SchemaDatabases: []string{"postgresql"},
		Gorm: struct {
			MysqDSN       string
			PostgresqlDSN string
		}{
			"root:password@tcp(127.0.0.1:3500)/go-rest-echo111?charset=utf8mb4&parseTime=True&loc=Local",
			"user=root password=password dbname=go-rest-echo111 host=127.0.0.1 port=5300 sslmode=disable TimeZone=Asia/Jakarta",
		},
		Monggo: struct {
			URI      string
			Database string
		}{
			"mongodb://root:password@localhost:27017/?readPreference=primary&ssl=false",
			"go-rest-echo",
		},
	}

	// test if error on connection
	want := new(db.Database)
	got, errs := db.NewDatabase(conf)

	assert.Equal(t, []error{errors.New("Can't connect database postgresql")}, errs)
	assert.Equal(t, want, got)

	// test if not error
	conf.Gorm.PostgresqlDSN = "user=root password=password dbname=go-rest-echo host=127.0.0.1 port=5300 sslmode=disable TimeZone=Asia/Jakarta"
	got, _ = db.NewDatabase(conf)

	assert.NotEqual(t, nil, got.Posrgresql)
}

func TestNewDatabase_ParamConfigSchemaDatabasesIsMongo(t *testing.T) {
	conf := &config.Config{
		SchemaDatabases: []string{"mongo"},
		Monggo: struct {
			URI      string
			Database string
		}{
			"q",
			"go-rest-echo1111",
		},
	}

	// test if error on connection
	want := new(db.Database)
	got, errs := db.NewDatabase(conf)

	assert.Equal(t, []error{errors.New("Can't connect database mongo")}, errs)
	assert.Equal(t, want, got)

	// test if not error
	conf.Monggo.URI = "mongodb://root:password@localhost:27018/?readPreference=primary&ssl=false"
	got, _ = db.NewDatabase(conf)

	assert.NotEqual(t, nil, got.Mongo)
}

func TestNewDatabase_ParamConfigSchemaDatabasesIsNotSupport(t *testing.T) {
	conf := &config.Config{
		SchemaDatabases: []string{"no-support-driver"},
	}

	// test if error on connection
	want := new(db.Database)
	got, errs := db.NewDatabase(conf)

	assert.Equal(t, []error{fmt.Errorf("schema database 'no-support-driver' is not support")}, errs)
	assert.Equal(t, want, got)
}
