package db

import (
	"errors"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var pgCon *gorm.DB

func postgresqlConnection() (err error) {
	pgCon, err = gorm.Open(postgres.Open(os.Getenv("GORM_POSTGRESQL_DSN")), &gorm.Config{})
	if err != nil {
		return errors.New("Can't connect database postgres with gorm library")
	}

	sqlCon, err := pgCon.DB()
	if err != nil {
		return errors.New("Can't instance postgres pool database connection")
	}

	sqlCon.SetMaxIdleConns(5)
	sqlCon.SetMaxOpenConns(50)

	return nil
}

// GetPostgresqlDB is
func GetPostgresqlDB() *gorm.DB {
	return pgCon
}
