package sentry

import (
	"go-service-echo/config"

	lib "github.com/getsentry/sentry-go"
)

// Sentry is
type Sentry struct {
	config *config.Config
}

// New is
func New(c *config.Config) *Sentry {
	if err := lib.Init(lib.ClientOptions{Dsn: c.Sentry.DNS, Environment: c.Sentry.ENV}); err != nil {
		return nil
	}

	return &Sentry{c}
}

// CaptureError is
func (s *Sentry) CaptureError(e error) *lib.EventID {
	return lib.CaptureException(e)
}

// CaptureMessage is
func (s *Sentry) CaptureMessage(msg string) *lib.EventID {
	return lib.CaptureMessage(msg)
}
