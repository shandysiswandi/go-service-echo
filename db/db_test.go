package db_test

import (
	"go-service-echo/config"
	"go-service-echo/db"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew_Error(t *testing.T) {
	is := assert.New(t)

	/* config is nil */
	actual, err := db.New(nil)
	is.Nil(actual)
	is.Equal(err, db.ErrDatabaseConfigIsNil)

	/* driver database is empty */
	actual, err = db.New(&config.DatabaseConfig{Driver: ""})
	is.Nil(actual)
	is.Equal(err, db.ErrNotUseDatabase)

	/* driver is mysql but error */
	actual, err = db.New(&config.DatabaseConfig{Driver: "mysql"})
	is.Nil(actual)
	is.Equal(err, db.ErrNotConnectMysql)

	/* driver is postgres but error */
	actual, err = db.New(&config.DatabaseConfig{Driver: "postgres"})
	is.Nil(actual)
	is.Equal(err, db.ErrNotConnectPostgres)

	/* driver is mongo but error */
	actual, err = db.New(&config.DatabaseConfig{Driver: "mongo"})
	is.Nil(actual)
	is.Equal(err, db.ErrNotConnectMongo)

	/* driver is not support */
	actual, err = db.New(&config.DatabaseConfig{Driver: "blabla"})
	is.Nil(actual)
	is.Equal(err, db.ErrDatabaseDriverNotSupport)
}

func TestNew_Success(t *testing.T) {
	is := assert.New(t)

	/* mysql config from env */
	os.Setenv("DB_DRIVER", "mysql")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_USERNAME", "root")
	os.Setenv("DB_PASSWORD", "password")
	os.Setenv("DB_NAME", "go-service")
	os.Setenv("DB_TIMEZONE", "UTC")

	actual, err := db.New(config.New().Database)
	is.NotNil(actual)
	is.NotNil(actual.SQL)
	is.Nil(actual.Mongo)
	is.Nil(err)

	/* postgres config from env */
	os.Setenv("DB_DRIVER", "postgres")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USERNAME", "postgres")
	os.Setenv("DB_PASSWORD", "password")
	os.Setenv("DB_NAME", "go-service")
	os.Setenv("DB_TIMEZONE", "UTC")

	actual, err = db.New(config.New().Database)
	is.NotNil(actual)
	is.NotNil(actual.SQL)
	is.Nil(actual.Mongo)
	is.Nil(err)

	/* mongo config from env */
	// os.Setenv("DB_DRIVER", "mongo")
	// os.Setenv("DB_HOST", "localhost")
	// os.Setenv("DB_PORT", "27017")
	// os.Setenv("DB_USERNAME", "root")
	// os.Setenv("DB_PASSWORD", "password")
	// os.Setenv("DB_NAME", "go-service")
	// os.Setenv("DB_TIMEZONE", "UTC")

	// actual, err = db.New(config.New().Database)
	// is.NotNil(actual)
	// is.NotNil(actual.SQL)
	// is.Nil(actual.Mongo)
	// is.Nil(err)
}
