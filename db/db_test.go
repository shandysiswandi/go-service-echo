package db_test

import (
	"go-rest-echo/config"
	"go-rest-echo/db"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestNew_Error(t *testing.T) {
	is := assert.New(t)

	ts := []struct {
		name     string
		dc       *config.DatabaseConfig
		expected error
	}{
		{"config is nil", nil, db.ErrConfigIsNil},
		{"driver is empty", &config.DatabaseConfig{Driver: ""}, db.ErrNotUseDatabase},
		{"driver mysql but error", &config.DatabaseConfig{Driver: "mysql"}, db.ErrNotConnectMysql},
		{"driver postgresql but error", &config.DatabaseConfig{Driver: "postgresql"}, db.ErrNotConnectPostgresql},
		{"driver mongo but error", &config.DatabaseConfig{Driver: "mongo"}, db.ErrNotConnectMongo},
		{"driver not support", &config.DatabaseConfig{Driver: "blabla"}, db.ErrDriverNotSupport},
	}

	for _, tc := range ts {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := db.New(tc.dc, "UTC")
			is.Nil(actual, tc.name)
			is.Equal(err, tc.expected, tc.name)
		})
	}
}

func TestNew_Success(t *testing.T) {
	is := assert.New(t)

	ts := []struct {
		name string
		file string
	}{
		{"mysql config from env", ".mysql"},
		{"postgresql config from env", ".postgresql"},
		{"mongo config from env", ".mongo"},
	}

	for _, tc := range ts {
		t.Run(tc.name, func(t *testing.T) {
			if err := godotenv.Overload(tc.file); err != nil {
				is.Nil(err)
			}

			dbConfig := config.New().Database
			actual, err := db.New(dbConfig, "UTC")

			if dbConfig.Driver == "mongo" {
				is.NotNil(actual)
				is.NotNil(actual.Mongo)
				is.Nil(actual.SQL)
				is.Nil(err)
			} else {
				is.NotNil(actual)
				is.NotNil(actual.SQL)
				is.Nil(actual.Mongo)
				is.Nil(err)
			}
		})
	}
}
