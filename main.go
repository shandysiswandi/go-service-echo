package main

import (
	"go-rest-echo/app"
	"go-rest-echo/config"
	"go-rest-echo/db"
	"go-rest-echo/util/logger"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		print(err)
	}

	config := config.New()

	db, err := db.New(config.Database, config.App.Timezone)
	if err != nil {
		logger.LogError(err)
	}

	app.New(config, db)
}
