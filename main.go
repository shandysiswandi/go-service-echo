package main

import (
	"go-rest-echo/app"
	"go-rest-echo/config"
	"go-rest-echo/db"
)

// Version and Build is
const (
	Version = "1.0.0"
	Build   = "0.0.1"
)

func main() {
	config.NewConfiguration()

	db.NewDatabase()

	app.NewApplication()
}
