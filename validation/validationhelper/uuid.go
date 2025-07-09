// Package validationhelper provides validation helper functions for govalid.
package validationhelper

// IsValidUUID validates UUID format manually for maximum performance.
// Validates RFC 4122 format: 8-4-4-4-12 hex digits with hyphens.
func IsValidUUID(s string) bool {
	// Check length: 36 characters (32 hex + 4 hyphens)
	if len(s) != 36 {
		return false
	}

	// Check hyphen positions: 8-4-4-4-12
	if !hasValidHyphens(s) {
		return false
	}

	// Check hex characters and version/variant
	if !hasValidHexChars(s) {
		return false
	}

	return isValidUUIDVersionAndVariant(s)
}

// isValidHexChar checks if a character is a valid hexadecimal digit.
func isValidHexChar(c byte) bool {
	return (c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')
}

// hasValidHyphens checks if hyphens are in correct positions.
func hasValidHyphens(s string) bool {
	return s[8] == '-' && s[13] == '-' && s[18] == '-' && s[23] == '-'
}

// hasValidHexChars checks if all non-hyphen characters are valid hex.
func hasValidHexChars(s string) bool {
	for i := 0; i < 36; i++ {
		if i == 8 || i == 13 || i == 18 || i == 23 {
			continue // skip hyphens
		}

		if !isValidHexChar(s[i]) {
			return false
		}
	}

	return true
}

// isValidUUIDVersionAndVariant checks version and variant fields.
func isValidUUIDVersionAndVariant(s string) bool {
	// Check version (position 14): must be 1-5
	version := s[14]
	if version < '1' || version > '5' {
		return false
	}

	// Check variant (position 19): must be 8, 9, A, B (case insensitive)
	variant := s[19]

	return variant == '8' || variant == '9' ||
		variant == 'A' || variant == 'a' ||
		variant == 'B' || variant == 'b'
}