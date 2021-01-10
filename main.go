package main

import (
	"log"

	"go-rest-echo/app/context"
	"go-rest-echo/app/middleware"
	"go-rest-echo/app/route"
	"go-rest-echo/app/validation"
	"go-rest-echo/database/mysql"
	"go-rest-echo/helper"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	/* --------------- Start Configuration --------------- */
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error %v", err)
	}

	/* --------------- Start Database --------------- */
	_, err = mysql.GetDB()
	if err != nil {
		log.Fatalf("Error %v", err)
	}

	// db.AutoMigrate(&entity.User{}, &entity.Task{})

	/* --------------- Start Aplication --------------- */
	e := echo.New()
	context.Initialize(e)
	validation.Initialize(e)
	middleware.Initialize(e)
	route.Initialize(e)
	e.Logger.Fatal(e.Start(":" + helper.Env("PORT")))
}
