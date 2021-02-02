package config

import "os"

type (
	// External is
	external struct {
		JsonplaceholderURL string
	}
)

func newExternalConfig() *external {
	return &external{
		JsonplaceholderURL: os.Getenv("EXTERTNAL_JSONPLACEHOLDER_URL"),
	}
}
