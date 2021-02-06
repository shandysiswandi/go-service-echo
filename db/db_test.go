package db_test

import (
	"go-rest-echo/config"
	"go-rest-echo/db"
	"log"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func init() {
	err := godotenv.Load(".env.test")
	if err != nil {
		log.Println(err)
		return
	}
}

func TestNewDatabase(t *testing.T) {
	var actual, errs = db.NewDatabase(config.NewConfiguration())

	assert.Equal(t, []error(nil), errs)

	assert.Equal(t, nil, actual.Mysql.Error)
	assert.Equal(t, int64(0), actual.Mysql.RowsAffected)

	assert.Equal(t, nil, actual.Postgresql.Error)
	assert.Equal(t, int64(0), actual.Postgresql.RowsAffected)

	assert.Equal(t, "go-rest-echo", actual.Mongo.Name())

}
