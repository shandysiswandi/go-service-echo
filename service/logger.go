package service

import (
	"go-rest-echo/config"

	log "github.com/sirupsen/logrus"
)

// Logger is
type Logger struct {
	config *config.Config
}

// NewLogger is
func NewLogger(c *config.Config) *Logger {
	return &Logger{config: c}
}

// Info is
func (logger *Logger) Info(val ...interface{}) error {
	log.Info(val)
	return nil
}

// Error is
func (logger *Logger) Error(val ...interface{}) error {
	log.Error(val)
	return nil
}
