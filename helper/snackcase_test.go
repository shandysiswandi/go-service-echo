package helper_test

import (
	"go-rest-echo/helper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSnakeCase(t *testing.T) {
	// instance assert
	assert := assert.New(t)

	// execution with closure or callback or step
	t.Run("1 All text is capitalize only one word", func(t *testing.T) {
		actual := helper.SnakeCase("CAPITALIZE")
		expected := "capitalize"

		assert.Equalf(actual, expected, "Expected `%s`, but actual %s", expected, actual)
	})

	t.Run("2 All text is lower with one word", func(t *testing.T) {
		actual := helper.SnakeCase("lower")
		expected := "lower"

		assert.Equalf(actual, expected, "Expected `%s`, but actual %s", expected, actual)
	})

	t.Run("3 Text input is IsUUID", func(t *testing.T) {
		actual := helper.SnakeCase("IsUUID")
		expected := "is_uuid"

		assert.Equalf(actual, expected, "Expected `%s`, but actual %s", expected, actual)
	})

	t.Run("4 Text input is UserId", func(t *testing.T) {
		actual := helper.SnakeCase("UserId")
		expected := "user_id"

		assert.Equalf(actual, expected, "Expected `%s`, but actual %s", expected, actual)
	})

	t.Run("5 Text input is theURL", func(t *testing.T) {
		actual := helper.SnakeCase("theURL")
		expected := "the_url"

		assert.Equalf(actual, expected, "Expected `%s`, but actual %s", expected, actual)
	})
}
