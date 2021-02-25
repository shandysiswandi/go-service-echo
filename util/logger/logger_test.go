package logger_test

import (
	"go-service-echo/util/logger"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInfo(t *testing.T) {
	assert.Nil(t, logger.Info("info"))
}

func TestError(t *testing.T) {
	assert.Nil(t, logger.Error("error"))
}
