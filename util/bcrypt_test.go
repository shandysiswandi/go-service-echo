package util_test

import (
	"go-rest-echo/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	// instance assert
	assert := assert.New(t)

	t.Run("1 Check five character from front no error", func(t *testing.T) {
		actual, err := util.HashPassword("password")
		expected := "$2a$10"

		assert.Nil(err)
		assert.Equalf(actual[0:6], expected, "Expected `%s`, but actual %s", expected, actual)
	})
}

func TestCheckPasswordHash(t *testing.T) {
	// instance assert
	assert := assert.New(t)

	t.Run("1 Check password is valid", func(t *testing.T) {
		actual := util.CheckPasswordHash("password", "$2a$10$mjqqoczR7odoHg/npdnwcuJCk4GHUDYTrkX48vuy/tNq7P/V/wAGi")
		expected := true

		assert.Equalf(expected, actual, "Expected `%v`, but actual %v", expected, actual)
	})
}
