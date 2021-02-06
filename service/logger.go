package service

import "go-rest-echo/config"

// Logger is
type Logger struct {
	config *config.Config
}

// NewLogger is
func NewLogger(c *config.Config) *Logger {
	return &Logger{config: c}
}

// Info is
func (logger *Logger) Info() error {
	return nil
}

// Error is
func (logger *Logger) Error() error {
	return nil
}
