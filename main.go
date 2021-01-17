package main

import (
	"go-rest-echo/app"
	"go-rest-echo/config"
	"go-rest-echo/db"
)

func main() {
	config.NewConfiguration()

	db.NewDatabase()

	app.NewApplication()
}
