package validationhelper

// IsValidAlphanum checks if the string contains only alphanumeric characters using regex.
// An empty string is considered invalid.
func IsValidAlphanum(s string) bool {
	for _, r := range s {
		if (r < 'a' || r > 'z') &&
			(r < 'A' || r > 'Z') &&
			(r < '0' || r > '9') {
			return false
		}
	}

	return s != ""
}
