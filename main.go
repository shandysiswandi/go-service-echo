package main

import (
	"go-rest-echo/app"
	"go-rest-echo/config"
	"go-rest-echo/db"
	"log"

	"github.com/joho/godotenv"
)

// Version and Build is
const (
	Version = "1.0.0"
	Build   = "0.0.1"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}

	conf := config.NewConfiguration()

	db, errs := db.NewDatabase(conf)
	if errs != nil {
		for _, e := range errs {
			log.Println(e)
		}
	}

	err = app.NewApplication(conf, db)
	if err != nil {
		log.Println(err)
	}
}
