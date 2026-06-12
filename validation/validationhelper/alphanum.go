package validationhelper

// IsValidAlphanum checks if the string contains only alphanumeric characters
// (a-z, A-Z, 0-9). An empty string is considered invalid.
func IsValidAlphanum(s string) bool {
	for i := 0; i < len(s); i++ {
		c := s[i]
		if (c < 'a' || c > 'z') && (c < 'A' || c > 'Z') && (c < '0' || c > '9') {
			return false
		}
	}

	return s != ""
}
