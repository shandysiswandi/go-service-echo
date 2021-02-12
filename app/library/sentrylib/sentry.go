package sentrylib

import (
	"go-rest-echo/config"

	"github.com/getsentry/sentry-go"
)

// Sentry is
type Sentry struct {
	config *config.Config
}

// New is
func New(c *config.Config) *Sentry {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:         c.Service.SentryDSN,
		Environment: c.App.Env,
	})
	if err != nil {
		return nil
	}
	return &Sentry{c}
}

// CaptureError is
func (s *Sentry) CaptureError(e error) *sentry.EventID {
	return sentry.CaptureException(e)
}

// CaptureMessage is
func (s *Sentry) CaptureMessage(msg string) *sentry.EventID {
	return sentry.CaptureMessage(msg)
}
