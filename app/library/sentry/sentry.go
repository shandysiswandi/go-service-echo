package sentry

import (
	"go-service-echo/config"

	lib "github.com/getsentry/sentry-go"
)

// Sentry is
type Sentry struct {
	config *config.SentryConfig
}

// New is
func New(c *config.SentryConfig) (*Sentry, error) {
	if err := lib.Init(lib.ClientOptions{Dsn: c.DNS, Environment: c.ENV}); err != nil {
		return nil, err
	}

	return &Sentry{c}, nil
}

// Error is
func (s *Sentry) Error(e error) *lib.EventID {
	return lib.CaptureException(e)
}

// Message is
func (s *Sentry) Message(msg string) *lib.EventID {
	return lib.CaptureMessage(msg)
}
