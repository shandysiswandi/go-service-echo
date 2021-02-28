package stringy

import (
	"regexp"
	"strings"
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
	var (
		matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
		matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
	)
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
