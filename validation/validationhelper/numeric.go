// Package validationhelper provides validation helper functions for govalid.
package validationhelper

// IsNumeric checks if a string contains only digit characters (0–9).
func IsNumeric(s string) bool {
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return false
		}
	}

	return s != ""
}
