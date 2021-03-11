package main

import (
	"context"
	"go-service-echo/app"
	"go-service-echo/config"
	"go-service-echo/config/constant"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	/********** ********** ********** ********** ********** ********** ********** **********/
	/* Load environment
	/********** ********** ********** ********** ********** ********** ********** **********/
	if err := godotenv.Load(); err != nil {
		println("Err: ", err)
	}

	/********** ********** ********** ********** ********** ********** ********** **********/
	/* Define config variable
	/********** ********** ********** ********** ********** ********** ********** **********/
	config := config.New()

	/********** ********** ********** ********** ********** ********** ********** **********/
	/* Define server and app variable
	/********** ********** ********** ********** ********** ********** ********** **********/
	app := app.New(config)

	/********** ********** ********** ********** ********** ********** ********** **********/
	/* run application server with graceful shutdown
	/********** ********** ********** ********** ********** ********** ********** **********/
	engine := app.GetEngine()
	engine.HideBanner = true
	engine.Server.ReadTimeout = 30 * time.Second
	engine.Server.WriteTimeout = 30 * time.Second
	println("⇨ Running", config.App.Name, "Version", constant.Version, "Build", constant.Build)

	go func() {
		if config.App.Env == constant.Production {
			err := engine.StartTLS(":"+config.App.Port, config.SSL.Cert, config.SSL.Key)
			if err != nil && err != http.ErrServerClosed {
				engine.Logger.Fatal("shutting down the server TLS")
			}
			return
		}

		err := engine.Start(":" + config.App.Port)
		if err != nil && err != http.ErrServerClosed {
			engine.Logger.Fatal("shutting down the server")
		}
	}()

	/********** ********** ********** ********** ********** ********** ********** **********/
	/* Wait for interrupt signal to gracefully shutdown the server with a timeout of 5 seconds.
	/* Use a buffered channel to avoid missing signals as recommended for signal.Notify
	/********** ********** ********** ********** ********** ********** ********** **********/
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	println("💥 Shutdown server ...")
	if err := engine.Shutdown(ctx); err != nil {
		println("Err", err)
	}
	println("💯 Shutdown server done !")
}
