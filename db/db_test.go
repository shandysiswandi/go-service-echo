package db_test

import (
	"go-rest-echo/config"
	"go-rest-echo/db"
	"log"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

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
