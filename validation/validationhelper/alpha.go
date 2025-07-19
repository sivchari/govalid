package validationhelper

// IsValidAlpha checks if the string contains only alphabetic characters using regex.
// An empty string is considered invalid.
func IsValidAlpha(s string) bool {
	for _, c := range s {
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
			continue
		}
	}

	return false
}
