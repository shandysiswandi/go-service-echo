package main

import (
	"go-rest-echo/app"
	"go-rest-echo/config"
	"go-rest-echo/db"
	"go-rest-echo/external"
	"go-rest-echo/service"
	"go-rest-echo/util"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic("NO .env GUYS")
	}

	config := config.NewConfiguration()
	db, errs := db.NewDatabase(config)
	if errs != nil {
		for _, e := range errs {
			util.LogError(e)
		}
	}

	service := service.New(config)
	external := external.New(config)
	app.NewApplicationAndServe(config, db, service, external)
}
