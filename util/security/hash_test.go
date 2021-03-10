package security_test

import (
	"go-service-echo/util/security"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Hash_Bcrypt(t *testing.T) {
	actual, err := security.Bcrypt("password", 10)
	expected := "$2a$10"

	assert.Nil(t, err)
	assert.Equal(t, expected, actual[0:6])
}

func Test_Hash_IsValidBcrypt(t *testing.T) {
	actual := security.IsValidBcrypt("password", "$2a$10$mjqqoczR7odoHg/npdnwcuJCk4GHUDYTrkX48vuy/tNq7P/V/wAGi")
	expected := true

	assert.Equal(t, expected, actual)
}

func Test_Hash_MD5(t *testing.T) {
	assert.Equal(t, "0cc175b9c0f1b6a831c399e269772661", security.MD5("a"))
}

func Test_Hash_Password(t *testing.T) {
	actual := security.Password("password")
	expected := "$2a$10"

	assert.NotEqual(t, "", actual[0:6])
	assert.Equal(t, expected, actual[0:6])
}

func Test_Hash_IsValidPassword(t *testing.T) {
	actual := security.IsValidPassword("password", "$2a$10$8ZJInOaW92x0EB8sHk/Dn.ZqAuQXuCv.b15DbaGsR4.FDD2nIqgR2")
	expected := true

	assert.Equal(t, expected, actual)
}
