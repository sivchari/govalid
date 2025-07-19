package validationhelper

import "regexp"

// IsValidAlpha checks if the string contains only alphabetic characters using regex.
// An empty string is considered invalid.
func IsValidAlpha(s string) bool {
	if s == "" {
		return false
	}
	
	match, _ := regexp.MatchString(`^[A-Za-z]+$`, s)
	return match
}
