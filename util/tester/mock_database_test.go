package tester_test

import (
	"go-service-echo/util/tester"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMockDatabase(t *testing.T) {
	a, b := tester.NewMockDatabase()

	assert.NotNil(t, a)
	assert.NotNil(t, b)
}

func TestMockMysqlGormConnection(t *testing.T) {
	a, b := tester.MockMysqlGormConnection()

	assert.NotNil(t, a)
	assert.NotNil(t, b)
}
