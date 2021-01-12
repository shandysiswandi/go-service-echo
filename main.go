package main

import (
	"go-rest-echo/app"
	"go-rest-echo/config"
	"go-rest-echo/db"
)

func main() {
	config.NewConfiguration().Start()

	db.NewDatabase().Start()

	app.NewApplication().Start()
}
