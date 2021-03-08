package stringy_test

import (
	"go-service-echo/util/stringy"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SnakeCase(t *testing.T) {
	assert.Equal(t, "capitalize", stringy.SnakeCase("CAPITALIZE"))
	assert.Equal(t, "lower", stringy.SnakeCase("lower"))
	assert.Equal(t, "is_uuid", stringy.SnakeCase("IsUUID"))
	assert.Equal(t, "user_id", stringy.SnakeCase("UserId"))
	assert.Equal(t, "the_url", stringy.SnakeCase("theURL"))
}

func Test_Split(t *testing.T) {
	assert.Equal(t, []string{"a", "2"}, stringy.Split("a/2", "/"))
	assert.Equal(t, []string{"a", "2"}, stringy.Split("/a/2", "/"))
	assert.Equal(t, []string{"a", "2"}, stringy.Split("/a/2/", "/"))
}

func Test_RandomString(t *testing.T) {
	assert.Equal(t, 10, len(stringy.RandomString(10)))
	assert.Equal(t, 10, len(stringy.RandomString(10)))
	assert.Equal(t, 50, len(stringy.RandomString(50)))
	assert.Equal(t, 100, len(stringy.RandomString(100)))
}
