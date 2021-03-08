package arrays_test

import (
	"go-service-echo/util/arrays"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_InArray(t *testing.T) {
	list := []string{"1", "2", "3", "4", "5"}
	assert.Equal(t, true, arrays.InArray(list, "2"))
	assert.Equal(t, false, arrays.InArray(list, "0"))
}
