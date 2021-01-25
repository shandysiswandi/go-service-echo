package db

import (
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func postgresqlConnection(dsn string) (db *gorm.DB, err error) {
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlCon, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlCon.SetMaxIdleConns(10 / 2)
	sqlCon.SetMaxOpenConns(100 / 2)
	sqlCon.SetConnMaxLifetime(time.Hour / 2)

	return db, nil
}
