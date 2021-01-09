package main

import (
	"log"

	"go-rest-echo/app"
	"go-rest-echo/database/mysql"
	"go-rest-echo/entity"
	"go-rest-echo/helper"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	config()
	db()
	start()
}

func config() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error %v", err)
	}
}

func db() {
	db, err := mysql.GetDB()
	if err != nil {
		log.Fatalf("Error %v", err)
	}

	db.AutoMigrate(&entity.User{}, &entity.Task{})
}

func start() {
	e := echo.New()
	app.Validation(e)
	app.Route(e)
	app.Middleware(e)
	e.Logger.Fatal(e.Start(":" + helper.Env("PORT")))
}
