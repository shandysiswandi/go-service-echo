package helper

import "os"

// Env is a function
func Env(env string) string {
	return os.Getenv(env)
}