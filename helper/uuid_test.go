package helper_test

import (
	"go-rest-echo/helper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateUUID(t *testing.T) {
	// instance assert
	assert := assert.New(t)

	actual := helper.GenerateUUID()

	assert.Equalf(36, len(actual), "Expected length = `%v`, but got `%v`", 36, len(actual))
}
