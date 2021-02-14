package main

import (
	"go-rest-echo/app"
	"go-rest-echo/config"
	"go-rest-echo/db"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		print(err)
	}

	config := config.New()

	db, err := db.NewDatabase(config.Database, config.App.Timezone)
	if err != nil {
		print(err)
	}

	app.New(config, db)
}
