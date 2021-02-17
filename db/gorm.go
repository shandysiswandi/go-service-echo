package db

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func gormConnection(dsn, driver string) (db *gorm.DB, err error) {
	var dialect gorm.Dialector
	var gormConfig *gorm.Config

	if driver == PostgresqlDriver {
		gormConfig = &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)}
		dialect = postgres.Open(dsn)
	} else {
		gormConfig = &gorm.Config{PrepareStmt: true, Logger: logger.Default.LogMode(logger.Silent)}
		dialect = mysql.Open(dsn)
	}

	db, err = gorm.Open(dialect, gormConfig)
	if err != nil {
		return nil, err
	}

	// pooling connection
	sqlCon, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlCon.SetMaxIdleConns(10 / 2)
	sqlCon.SetMaxOpenConns(100 / 2)
	sqlCon.SetConnMaxLifetime(time.Hour / 2)

	return db, nil
}
