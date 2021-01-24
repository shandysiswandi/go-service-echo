package config_test

import (
	"go-rest-echo/config"
	"log"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func init() {
	err := godotenv.Load("../.env.test")
	if err != nil {
		log.Println(err)
	}
}

func TestNewConfiguration(t *testing.T) {
	expected := &config.Config{
		App: struct {
			Env  string
			Port string
			Name string
		}{
			"env",
			"port",
			"name",
		},
		SchemaDatabases: []string{"SchemaDatabases", "SchemaDatabases", "SchemaDatabases"},
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

	got := config.NewConfiguration()
	assert.Equal(t, expected, got)
}
