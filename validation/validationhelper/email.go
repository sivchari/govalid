// Package validationhelper provides validation helper functions for govalid.
package validationhelper

// IsValidEmail validates email format manually for maximum performance.
// Validates basic email structure: local@domain format.
func IsValidEmail(email string) bool {
	// Basic length check
	if len(email) < 5 || len(email) > 254 { // a@b.c = 5 chars minimum, RFC 5321 limit
		return false
	}

	// Find the @ symbol position
	atIndex := findAtSymbol(email)
	if atIndex == -1 {
		return false
	}

	// Validate local and domain parts
	local := email[:atIndex]
	domain := email[atIndex+1:]

	return isValidLocalPart(local) && isValidDomainPart(domain)
}

// findAtSymbol finds the position of @ symbol and validates there's exactly one.
func findAtSymbol(email string) int {
	atIndex := -1
	atCount := 0

	for i, c := range email {
		if c == '@' {
			atIndex = i
			atCount++
		}
	}

	if atCount != 1 || atIndex <= 0 || atIndex >= len(email)-1 {
		return -1
	}

	return atIndex
}

// isValidLocalPart validates the local part (before @) of an email.
func isValidLocalPart(local string) bool {
	if len(local) == 0 || len(local) > 64 { // RFC 5321 limit
		return false
	}

	return isValidLocalPartFormat(local) && isValidLocalPartChars(local)
}

// isValidLocalPartFormat checks dot rules in local part.
func isValidLocalPartFormat(local string) bool {
	// Check for consecutive dots or leading/trailing dots
	if local[0] == '.' || local[len(local)-1] == '.' {
		return false
	}

	for i := 0; i < len(local)-1; i++ {
		if local[i] == '.' && local[i+1] == '.' {
			return false
		}
	}

	return true
}

// isValidLocalPartChars checks allowed characters in local part.
func isValidLocalPartChars(local string) bool {
	for _, c := range local {
		if !isValidLocalChar(c) {
			return false
		}
	}

	return true
}

// isValidLocalChar checks if a character is valid in local part.
func isValidLocalChar(c rune) bool {
	// Alphanumeric characters
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
		return true
	}

	return isValidLocalSpecialChar(c)
}

// isValidLocalSpecialChar checks if a character is a valid special character in local part.
func isValidLocalSpecialChar(c rune) bool {
	// RFC 5322 allowed special characters
	switch c {
	case '.', '_', '-', '+', '=', '!', '#', '$', '%', '&', '\'', '*', '/', '?', '^', 96, '{', '|', '}', '~':
		return true // c == 96 is backtick
	default:
		return false
	}
}

// isValidDomainPart validates the domain part (after @) of an email.
func isValidDomainPart(domain string) bool {
	if len(domain) == 0 || len(domain) > 253 { // RFC 1035 limit
		return false
	}

	// Must contain at least one dot
	hasDot := false

	for _, c := range domain {
		if c == '.' {
			hasDot = true

			break
		}
	}

	if !hasDot {
		return false
	}

	// Cannot start or end with dot or hyphen
	if domain[0] == '.' || domain[len(domain)-1] == '.' ||
		domain[0] == '-' || domain[len(domain)-1] == '-' {
		return false
	}

	return validateDomainLabels(domain)
}

// validateDomainLabels parses and validates each domain label.
func validateDomainLabels(domain string) bool {
	labelCount := 0
	start := 0

	for i := 0; i <= len(domain); i++ {
		if i == len(domain) || domain[i] == '.' {
			if i == start {
				// Empty label
				return false
			}

			label := domain[start:i]
			labelCount++

			if !isValidDomainLabel(label) {
				return false
			}

			start = i + 1
		}
	}

	// Must have at least 2 labels (e.g., "example.com")
	return labelCount >= 2
}

// isValidDomainLabel validates a single domain label.
func isValidDomainLabel(label string) bool {
	// Validate label length
	if len(label) > 63 { // RFC 1035 limit
		return false
	}

	// Cannot start or end with hyphen
	if label[0] == '-' || label[len(label)-1] == '-' {
		return false
	}

	return isValidDomainLabelChars(label)
}

// isValidDomainLabelChars checks if all characters in label are valid.
func isValidDomainLabelChars(label string) bool {
	for _, c := range label {
		if !isValidDomainChar(c) {
			return false
		}
	}

	return true
}

// isValidDomainChar checks if a character is valid in domain label.
func isValidDomainChar(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9') || c == '-'
}
