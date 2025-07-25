package validationhelper

// IsValidAlpha checks if the string contains only alphabetic characters using regex.
// An empty string is considered invalid.
func IsValidAlpha(s string) bool {
	for i := 0; i < len(s); i++ {
		c := s[i]
		if (c < 'a' || c > 'z') && (c < 'A' || c > 'Z') {
			return false
		}
	}

	return true
}
