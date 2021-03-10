package security

import (
	"crypto/md5"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Bcrypt is
func Bcrypt(plain string, salt int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(plain), salt)
	return string(bytes), err
}

// IsValidBcrypt is
func IsValidBcrypt(plain, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(plain)) == nil
}

// MD5 is
func MD5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// Password is
func Password(plain string) string {
	s := MD5(plain)
	s, err := Bcrypt(s, 10)
	if err != nil {
		return ""
	}
	return s
}

// IsValidPassword is
func IsValidPassword(plain, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(MD5(plain))) == nil
}
