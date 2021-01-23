package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func mysqlConnection(dsn string) (db *gorm.DB, err error) {
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{PrepareStmt: true})
	if err != nil {
		return nil, err
	}

	// pooling connection
	sqlCon, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlCon.SetMaxIdleConns(5)
	sqlCon.SetMaxOpenConns(50)

	return db, nil
}
