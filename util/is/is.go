package is

// InArray :
func InArray(arr []string, ins ...string) bool {
	if arr == nil {
		return false
	}

	if len(arr) < 1 {
		return false
	}

	for _, ar := range arr {
		for _, in := range ins {
			if ar == in {
				return true
			}
		}
	}

	return false
}
