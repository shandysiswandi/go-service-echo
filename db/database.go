package db

import (
	"go-rest-echo/db/mysql"
	"log"
)

// NewDatabase is
func NewDatabase() {
	err := mysql.GormMysqlConnection()
	if err != nil {
		log.Println(err.Error())
	}
}
