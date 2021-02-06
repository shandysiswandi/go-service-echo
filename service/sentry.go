package service

import "go-rest-echo/config"

// Sentry is
type Sentry struct {
	config *config.Config
}

// NewSentry is
func NewSentry(c *config.Config) *Sentry {
	return &Sentry{c}
}

// Log is
func (Sentry) Log() {
	//
}
