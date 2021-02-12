package external

import (
	"go-rest-echo/config"
	"go-rest-echo/external/jsonplaceholder.typicode.com"
)

type (
	// External is
	External struct {
		JSONPlaceHolder *jsonplaceholder.JSONPlaceHolder
	}
)

// New is
func New(c *config.Config) *External {
	if c == nil {
		return nil
	}

	return &External{
		JSONPlaceHolder: jsonplaceholder.New(c.External.JsonplaceholderURL),
	}
}
