package utils

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// CopyMap copy src to dest
func CopyMap(dest, src map[string]interface{}) {
	for k := range src {
		dest[k] = src[k]
	}
}
