package mysql

import (
	"errors"
	"fmt"
	"go-rest-echo/helper"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB is a instance to
var DB *gorm.DB

// GetDB is
func GetDB() (con *gorm.DB, err error) {
	var (
		user = helper.Env("GORM_USERNAME")
		pass = helper.Env("GORM_PASSWORD")
		host = helper.Env("GORM_HOST")
		port = helper.Env("GORM_PORT")
		db   = helper.Env("GORM_DATABASE")
		dsn  = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, db)
	)

	con, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{}),
	})
	if err != nil {
		err = errors.New("Can't connect database mysql with gorm library")
		return nil, err
	}

	conDB, err := con.DB()
	conDB.SetMaxIdleConns(5)
	conDB.SetMaxOpenConns(50)
	DB = con
	return con, nil
}
