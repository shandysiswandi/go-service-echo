package arrays_test

import (
	"go-service-echo/util/arrays"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Split(t *testing.T) {
	ts := []struct {
		s, sep   string
		expected []string
	}{
		{"", "", []string{}},
		{"abc", "", []string{}},
		{"", "_", []string{}},

		{"the/url", "/", []string{"the", "url"}},
		{"/the/url", "/", []string{"the", "url"}},
		{"/the/url/", "/", []string{"the", "url"}},

		{"/a /b /c", "/", []string{"a", "b", "c"}},
	}

	for _, val := range ts {
		assert.Equal(t, val.expected, arrays.Split(val.s, val.sep))
	}
}
