package bcrypt_test

import (
	"go-service-echo/util/bcrypt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HashPassword(t *testing.T) {
	is := assert.New(t)

	actual, err := bcrypt.HashPassword("password")
	expected := "$2a$10"

	is.Nil(err)
	is.Equal(expected, actual[0:6])
}

func Test_IsValidPassword(t *testing.T) {
	is := assert.New(t)

	actual := bcrypt.IsValidPassword("password", "$2a$10$mjqqoczR7odoHg/npdnwcuJCk4GHUDYTrkX48vuy/tNq7P/V/wAGi")
	expected := true

	is.Equal(expected, actual)
}
