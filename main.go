package main

import (
	"go-service-echo/app"
	"go-service-echo/config"
	"go-service-echo/db"
	"go-service-echo/util/logger"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		logger.Error(err)
	}

	config := config.New()

	db, err := db.New(config.Database)
	if err != nil {
		logger.Error(err)
	}

	app.New(config, db)
}
