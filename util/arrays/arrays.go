package arrays

// InArray is
func InArray(a []string, b string) bool {
	for _, a := range a {
		if a == b {
			return true
		}
	}
	return false
}
