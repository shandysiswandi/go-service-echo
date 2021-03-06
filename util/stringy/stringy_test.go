package stringy_test

import (
	"go-service-echo/util/stringy"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SnakeCase(t *testing.T) {
	ts := []struct {
		input    string
		expected string
	}{
		{"CAPITALIZE", "capitalize"},
		{"lower", "lower"},
		{"IsUUID", "is_uuid"},
		{"UserId", "user_id"},
		{"theURL", "the_url"},
	}

	for _, val := range ts {
		assert.Equal(t, val.expected, stringy.SnakeCase(val.input))
	}
}
