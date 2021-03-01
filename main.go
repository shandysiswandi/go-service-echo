package main

import (
	"context"
	"go-service-echo/app"
	"go-service-echo/config"
	"go-service-echo/config/constant"
	"go-service-echo/util/logger"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
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
		RegisterValidation().
		RegisterMiddlewares().
		RegisterRoutes()

	/********** ********** ********** **********/
	/* run application server with graceful shutdown
	/********** ********** ********** **********/
	engine := app.GetEngine()
	go server(config, engine)

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 5 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := engine.Shutdown(ctx); err != nil {
		engine.Logger.Fatal(err)
	}
}

func server(config *config.Config, app *echo.Echo) {
	if config.App.Env == constant.Production {
		err := app.StartTLS(":"+config.App.Port, config.SSL.Cert, config.SSL.Key)
		if err != nil && err != http.ErrServerClosed {
			app.Logger.Fatal("shutting down the server TLS")
		}
		return
	}

	err := app.Start(":" + config.App.Port)
	if err != nil && err != http.ErrServerClosed {
		app.Logger.Fatal("shutting down the server")
	}
}
