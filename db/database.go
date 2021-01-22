package db

import (
	"log"
)

// NewDatabase is
func NewDatabase() {
	var err error

	err = mysqlConnection()
	if err != nil {
		log.Println(err.Error())
	}

	err = postgresqlConnection()
	if err != nil {
		log.Println(err.Error())
	}

	// err := mongoConnection()
	// if err != nil {
	// 	log.Println(err.Error())
	// }
}
