package db

import (
	"errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func postgresqlConnection(dsn string) (db *gorm.DB, err error) {
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.New("Can't connect database postgres with gorm library")
	}

	sqlCon, err := db.DB()
	if err != nil {
		return nil, errors.New("Can't instance postgres pool database connection")
	}

	sqlCon.SetMaxIdleConns(5)
	sqlCon.SetMaxOpenConns(50)

	return db, nil
}
