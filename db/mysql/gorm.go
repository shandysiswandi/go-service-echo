package mysql

import (
	"errors"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var connection *gorm.DB

// GormMysqlConnection is
func GormMysqlConnection() (err error) {
	var dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("GORM_USERNAME"),
		os.Getenv("GORM_PASSWORD"),
		os.Getenv("GORM_HOST"),
		os.Getenv("GORM_PORT"),
		os.Getenv("GORM_DATABASE"),
	)

	connection, err = gorm.Open(mysql.Open(dsn), &gorm.Config{PrepareStmt: true})
	if err != nil {
		return errors.New("Can't connect database mysql with gorm library")
	}

	sqlCon, err := connection.DB()
	sqlCon.SetMaxIdleConns(5)
	sqlCon.SetMaxOpenConns(50)

	return nil
}

// GetDB is
func GetDB() *gorm.DB {
	return connection
}
