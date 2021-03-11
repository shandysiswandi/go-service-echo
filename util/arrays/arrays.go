package arrays

import (
	"strings"
)

// Split is
func Split(s, sep string) []string {
	temp := make([]string, 0)
	if s == "" || sep == "" {
		return temp
	}
	for _, s = range strings.Split(s, sep) {
		s = strings.TrimSpace(s)
		if s != "" {
			temp = append(temp, s)
		}
	}
	return temp
}
