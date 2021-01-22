package helper_test

import (
	"go-rest-echo/helper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	// instance assert
	assert := assert.New(t)

	t.Run("1 Check five character from front no error", func(t *testing.T) {
		actual, err := helper.HashPassword("password")
		expected := "$2a$10"

		assert.Nilf(err, "Expected `%s`, but got error", expected)
		assert.Equalf(actual[0:6], expected, "Expected `%s`, but actual %s", expected, actual)
	})
}

func TestCheckPasswordHash(t *testing.T) {
	// instance assert
	assert := assert.New(t)

	t.Run("1 Check password is valid", func(t *testing.T) {
		actual := helper.CheckPasswordHash("password", "$2a$10$mjqqoczR7odoHg/npdnwcuJCk4GHUDYTrkX48vuy/tNq7P/V/wAGi")
		expected := true

		assert.Equalf(expected, actual, "Expected `%v`, but actual %v", expected, actual)
	})
}
