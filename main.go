package main

import (
	"go-service-echo/app"
	"go-service-echo/config"
	"go-service-echo/config/constant"
	"go-service-echo/util/logger"

	"github.com/joho/godotenv"
)

func main() {
	/********** ********** ********** **********/
	/* Load environment
	/********** ********** ********** **********/
	if err := godotenv.Load(); err != nil {
		logger.Error(err)
	}

	/********** ********** ********** **********/
	/* Define config variable
	/********** ********** ********** **********/
	config := config.New()

	/********** ********** ********** **********/
	/* Define server and app variable
	/********** ********** ********** **********/
	app := app.New(config).
		RegisterContext().
		SetValidation().
		RegisterMiddlewares().
		RegisterRoutes()

	/********** ********** ********** **********/
	/* run application server
	/********** ********** ********** **********/
	appLogger := app.GetLogger()
	appEngine := app.GetEngine()

	if config.App.Env == constant.Production {
		appLogger.Fatal(appEngine.StartTLS(":"+config.App.Port, config.SSL.Cert, config.SSL.Key))
	}

	appLogger.Fatal(appEngine.Start(":" + config.App.Port))
}
