package is_test

import (
	"go-service-echo/util/is"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_InArray(t *testing.T) {
	assert.Equal(t, false, is.InArray(nil, ""))
	assert.Equal(t, false, is.InArray([]string{}, ""))
	assert.Equal(t, true, is.InArray([]string{"", ""}, ""))
	assert.Equal(t, false, is.InArray([]string{"a", "b"}, ""))
}
