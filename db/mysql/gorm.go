package mysql

import (
	"errors"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var connection *gorm.DB

type database struct {
	dsn string
}

// Database is
type Database interface {
	Initialize() error
}

// NewDatabase is
func NewDatabase() Database {
	return &database{
		dsn: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			os.Getenv("GORM_USERNAME"),
			os.Getenv("GORM_PASSWORD"),
			os.Getenv("GORM_HOST"),
			os.Getenv("GORM_PORT"),
			os.Getenv("GORM_DATABASE"),
		),
	}
}

func (d *database) Initialize() (err error) {
	connection, err = gorm.Open(mysql.Open(d.dsn), &gorm.Config{PrepareStmt: true})
	if err != nil {
		err = errors.New("Can't connect database mysql with gorm library")
		return err
	}

	conDB, err := connection.DB()
	conDB.SetMaxIdleConns(5)
	conDB.SetMaxOpenConns(50)
	return nil
}

// GetDB is
func GetDB() *gorm.DB {
	return connection
}
