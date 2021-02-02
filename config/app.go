package config

import "os"

// App is struct
type app struct {
	Env  string
	Port string
	Name string
}

func newAppConfig() *app {
	return &app{
		Env:  os.Getenv("APP_ENV"),
		Port: os.Getenv("APP_PORT"),
		Name: os.Getenv("APP_NAME"),
	}
}
