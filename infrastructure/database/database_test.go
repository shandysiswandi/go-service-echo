package database_test

import (
	"go-service-echo/config"
	"go-service-echo/infrastructure/database"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew_Error(t *testing.T) {
	is := assert.New(t)

	/* config is nil */
	actual, err := database.New(nil)
	is.Nil(actual)
	is.Equal(err, database.ErrDatabaseConfigIsNil)

	/* driver database is empty */
	actual, err = database.New(&config.DatabaseConfig{Driver: ""})
	is.Nil(actual)
	is.Equal(err, database.ErrNotUseDatabase)

	/* driver is mysql but error */
	actual, err = database.New(&config.DatabaseConfig{Driver: "mysql"})
	is.Nil(actual)
	is.Equal(err, database.ErrNotConnectMysql)

	/* driver is postgres but error */
	actual, err = database.New(&config.DatabaseConfig{Driver: "postgres"})
	is.Nil(actual)
	is.Equal(err, database.ErrNotConnectPostgres)

	/* driver is mongo but error */
	actual, err = database.New(&config.DatabaseConfig{Driver: "mongo"})
	is.Nil(actual)
	is.Equal(err, database.ErrNotConnectMongo)

	/* driver is not support */
	actual, err = database.New(&config.DatabaseConfig{Driver: "blabla"})
	is.Nil(actual)
	is.Equal(err, database.ErrDatabaseDriverNotSupport)
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

	actual, err := database.New(config.New().Database)
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

	actual, err = database.New(config.New().Database)
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

	// actual, err = database.New(config.New().Database)
	// is.NotNil(actual)
	// is.NotNil(actual.SQL)
	// is.Nil(actual.Mongo)
	// is.Nil(err)
}
