package main

import (
	"go-rest-echo/app"
	"go-rest-echo/config"
	"go-rest-echo/db"
	"go-rest-echo/util"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic("NO .env GUYS")
	}

	config := config.New()
	db, errs := db.NewDatabase(config)
	if errs != nil {
		for _, e := range errs {
			util.LogError(e)
		}
	}

	app.NewApplicationAndServe(config, db)
}
