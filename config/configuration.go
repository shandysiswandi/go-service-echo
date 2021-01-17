package config

import "github.com/joho/godotenv"

// NewConfiguration is
func NewConfiguration() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}
