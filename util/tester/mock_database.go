package tester

import (
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewMockDatabase is
func NewMockDatabase() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		panic("an error " + err.Error() + " was not expected when opening a stub database connection")
	}

	return db, mock
}

// MockMysqlGormConnection is
func MockMysqlGormConnection() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock := NewMockDatabase()
	mysqlConfig := mysql.Config{Conn: db, SkipInitializeWithVersion: true}
	gormConfig := &gorm.Config{Logger: nil}

	gorm, err := gorm.Open(mysql.New(mysqlConfig), gormConfig)
	if err != nil {
		panic("an error " + err.Error() + " was not expected when opening a mysql gorm database connection")
	}

	return gorm, mock
}
