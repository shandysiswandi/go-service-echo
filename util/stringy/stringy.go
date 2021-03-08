package stringy

import (
	"math/rand"
	"regexp"
	"strings"
	"time"
)

const (
	charsetLower  = "abcdefghijklmnopqrstuvwxyz"
	charsetUpper  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charsetNumber = "0123456789"
)

// Split is
func Split(a, b string) []string {
	re := make([]string, 0)
	for _, a = range strings.Split(a, b) {
		if strings.TrimSpace(a) != "" {
			re = append(re, a)
		}
	}
	return re
}

// SnakeCase is
func SnakeCase(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

// RandomString is
func RandomString(l int) string {
	b := make([]byte, l)
	for i := range b {
		b[i] = charsetLower[rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(charsetLower))]
	}
	return string(b)
}
