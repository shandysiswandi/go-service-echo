package main

import (
	"go-service-echo/app"
	"go-service-echo/config"
	"go-service-echo/db"
	"go-service-echo/util/logger"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	/********** ********** ********** **********/
	/* Load environment
	/********** ********** ********** **********/
	if err := godotenv.Load(".env"); err != nil {
		logger.Error(err)
	}

	/********** ********** ********** **********/
	/* Define config variable
	/********** ********** ********** **********/
	config := config.New()

	/********** ********** ********** **********/
	/* Define database variable
	/********** ********** ********** **********/
	db, err := db.New(config.Database)
	if err != nil {
		logger.Error(err)
	}

	/********** ********** ********** **********/
	/* Define server and app variable then run it
	/********** ********** ********** **********/
	app.New(echo.New(), config, db).
		SetContext().
		SetValidation().
		SetMiddlewares().
		SetRoutes().
		Run()
}
