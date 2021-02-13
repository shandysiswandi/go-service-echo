package util_test

import (
	"go-rest-echo/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogInfo(t *testing.T) {
	assert.Nil(t, util.LogInfo("info"))
}

func TestLogError(t *testing.T) {
	assert.Nil(t, util.LogError("error"))
}
