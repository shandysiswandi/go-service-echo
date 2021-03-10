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

func Test_Random(t *testing.T) {
	ts := []struct {
		input    int
		expected int
	}{
		{5, 5},
		{10, 10},
		{50, 50},
		{100, 100},
		{150, 150},
	}

	for _, val := range ts {
		assert.Equal(t, val.expected, len(stringy.Random(val.input)))
	}
}
