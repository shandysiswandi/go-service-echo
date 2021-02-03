package main

import (
	"go-rest-echo/app"
	"go-rest-echo/config"
	"go-rest-echo/db"
	"go-rest-echo/external"
	"go-rest-echo/service"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}
}

func main() {
	conf := config.NewConfiguration()

	db, errs := db.NewDatabase(conf)
	if errs != nil {
		for _, e := range errs {
			log.Println(e)
		}
	}

	serv := service.New(conf)

	ext := external.New()

	app.NewApplicationAndServe(conf, db, serv, ext)
}
