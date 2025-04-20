package tools

// Convert an string to a rune sequence
func ToRuneArray(s string) []rune {
	result := []rune{}
	for _, char := range s {
		result = append(result, char)
	}
	return result
}

// Compare two strings. Return 1 if a > b(alphabetically), -1 if a < b and 0 in other case
// if one is prefix of the other, them the shortest is taken as the greater
func CompareString(s1 string, s2 string) int {
	min := len(s1)
	if len(s2) < min {
		min = len(s2)
	}
	for i := 0; i < min; i++ {
		if s1[i] < s2[i] {
			return -1
		}
		if s1[i] > s2[i] {
			return 1
		}
	}
	if len(s1) < len(s2) {
		return 1
	}
	if len(s1) > len(s2) {
		return -1
	}
	return 0
}
