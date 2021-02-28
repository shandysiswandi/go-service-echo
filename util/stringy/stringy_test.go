package stringy_test

import (
	"go-service-echo/util/stringy"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSnakeCase(t *testing.T) {
	is := assert.New(t)

	cases := []struct {
		desc     string
		input    string
		expected string
	}{
		{"All text is capitalize only one word", "CAPITALIZE", "capitalize"},
		{"All text is lower with one word", "lower", "lower"},
		{"Text input is IsUUID", "IsUUID", "is_uuid"},
		{"Text input is UserId", "UserId", "user_id"},
		{"Text input is theURL", "theURL", "the_url"},
	}

	for _, tc := range cases {
		is.Equal(tc.expected, stringy.SnakeCase(tc.input))
	}
}

func TestSplit(t *testing.T) {
	is := assert.New(t)

	cases := []struct {
		desc     string
		inA, inB string
		expected []string
	}{
		{"test 1", "a/2", "/", []string{"a", "2"}},
		{"test 2", "/a/2", "/", []string{"a", "2"}},
		{"test 3", "/a/2/", "/", []string{"a", "2"}},
	}

	for _, tc := range cases {
		is.Equal(tc.expected, stringy.Split(tc.inA, tc.inB))
	}
}
