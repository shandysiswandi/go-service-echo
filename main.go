package main

import (
	"go-rest-echo/app"
	"go-rest-echo/config"
	"go-rest-echo/db"
	"log"
)

// Version and Build is
const (
	Version = "1.0.0"
	Build   = "0.0.1"
)

func main() {
	conf, err := config.NewConfiguration()
	if err != nil {
		panic(err)
	}

	db, errs := db.NewDatabase(conf)
	if errs != nil {
		for _, e := range errs {
			log.Println(e)
		}
	}

	err = app.NewApplication(conf, db)
	if err != nil {
		panic(err)
	}
}
