package tools

// Convert an string to a rune sequence
func ToRuneArray(s string) []rune {
	result := []rune{}
	for _, char := range s {
		result = append(result, char)
	}
	return result
}
