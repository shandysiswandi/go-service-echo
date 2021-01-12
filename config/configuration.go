package config

import "github.com/joho/godotenv"

type config struct{}

// Interface is
type Interface interface {
	Start()
}

// NewConfiguration is
func NewConfiguration() Interface {
	return &config{}
}

func (config) Start() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}
