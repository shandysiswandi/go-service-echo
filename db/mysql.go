package db

import (
	"errors"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var mysqlCon *gorm.DB

func mysqlConnection() (err error) {
	mysqlCon, err = gorm.Open(mysql.Open(os.Getenv("GORM_MYSQL_DSN")), &gorm.Config{PrepareStmt: true})
	if err != nil {
		return errors.New("Can't connect database mysql with gorm library")
	}

	// pooling connection
	sqlCon, err := mysqlCon.DB()
	if err != nil {
		return errors.New("Can't instance mysql pool database connection")
	}

	sqlCon.SetMaxIdleConns(5)
	sqlCon.SetMaxOpenConns(50)

	return nil
}

// GetMysqlDB is
func GetMysqlDB() *gorm.DB {
	return mysqlCon
}
