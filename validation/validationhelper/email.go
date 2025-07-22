// Package validationhelper provides validation helper functions for govalid.
package validationhelper

// IsValidEmail validates email format manually for maximum performance.
// This function implements a simplified but practical email validation algorithm based on
// RFC 5321 (SMTP) and RFC 5322 (Internet Message Format) specifications.
//
// Email format: local-part@domain
//
// Validation steps:
//  1. Overall length check (5-254 characters per RFC 5321)
//  2. Find and validate exactly one @ symbol
//  3. Validate local part (before @):
//     - Length: 1-64 characters (RFC 5321)
//     - No leading/trailing dots
//     - No consecutive dots
//     - Allowed characters: a-z, A-Z, 0-9, and special chars: .!#$%&'*+-/=?^_`{|}~
//  4. Validate domain part (after @):
//     - Length: 1-253 characters (RFC 1035)
//     - Must contain at least one dot
//     - Must have at least 2 labels (e.g., "example.com")
//     - Each label: 1-63 characters, alphanumeric + hyphen, no leading/trailing hyphen
//
// Example breakdown:
//
//	user.name+tag@sub.example.com
//	^^^^^^^^^|^^^|^^^^^^^^^^^^^^^^
//	local    |tag|domain
//	         +   @ (separator)
//
// Note: This implementation does not support:
//   - Quoted strings in local part (e.g., "john doe"@example.com)
//   - IP addresses as domains (e.g., user@[192.168.1.1])
//   - Internationalized domain names (IDN)
//   - Comments in email addresses
func IsValidEmail(email string) bool {
	// Step 1: Basic length check
	// Minimum: a@b.c = 5 characters
	// Maximum: 254 characters (RFC 5321 section 4.5.3.1.3)
	if len(email) < 5 || len(email) > 254 {
		return false
	}

	// Step 2: Find the @ symbol position
	// Must have exactly one @ symbol, not at the beginning or end
	atIndex := findAtSymbol(email)
	if atIndex == -1 {
		return false
	}

	// Step 3 & 4: Validate local and domain parts
	// Split email at @ symbol
	local := email[:atIndex]    // Everything before @
	domain := email[atIndex+1:] // Everything after @

	return isValidLocalPart(local) && isValidDomainPart(domain)
}

// findAtSymbol finds the position of @ symbol and validates there's exactly one.
// Returns -1 if:
//   - No @ symbol found
//   - Multiple @ symbols found
//   - @ is at the beginning (position 0)
//   - @ is at the end (last position)
//
// Example:
//
//	"user@domain.com" -> returns 4
//	"@domain.com"     -> returns -1 (at beginning)
//	"user@"           -> returns -1 (at end)
//	"user@@domain"    -> returns -1 (multiple @)
func findAtSymbol(email string) int {
	atIndex := -1
	atCount := 0

	for i, c := range email {
		if c == '@' {
			atIndex = i
			atCount++
		}
	}

	// Validate: exactly one @, not at beginning or end
	// atIndex <= 0: @ is at position 0 or not found
	// atIndex >= len(email)-1: @ is at last position
	if atCount != 1 || atIndex <= 0 || atIndex >= len(email)-1 {
		return -1
	}

	return atIndex
}

// isValidLocalPart validates the local part (before @) of an email.
// RFC 5321 section 4.5.3.1.1 specifies maximum length of 64 octets.
//
// Valid examples:
//
//	"user"           -> simple username
//	"user.name"      -> with dot
//	"user+tag"       -> with plus addressing
//	"user_name-123"  -> with underscore and hyphen
//
// Invalid examples:
//
//	".user"          -> starts with dot
//	"user."          -> ends with dot
//	"user..name"     -> consecutive dots
//	"" (empty)       -> no characters
func isValidLocalPart(local string) bool {
	// Length check: 1-64 characters (RFC 5321 section 4.5.3.1.1)
	if local == "" || len(local) > 64 {
		return false
	}

	return isValidLocalPartFormat(local) && isValidLocalPartChars(local)
}

// isValidLocalPartFormat checks dot rules in local part.
// RFC 5322 section 3.4.1: dots cannot appear consecutively or at beginning/end.
//
// Examples with positions:
//
//	"user.name"  -> valid
//	0123456789
//
//	".username"  -> invalid (dot at position 0)
//	0123456789
//
//	"user..name" -> invalid (consecutive dots at positions 4,5)
//	01234567890
func isValidLocalPartFormat(local string) bool {
	// Check for leading or trailing dots
	if local[0] == '.' || local[len(local)-1] == '.' {
		return false
	}

	// Check for consecutive dots
	// Iterate through string checking each character and the next
	for i := 0; i < len(local)-1; i++ {
		if local[i] == '.' && local[i+1] == '.' {
			return false
		}
	}

	return true
}

// isValidLocalPartChars checks allowed characters in local part.
// Each character must be validated against RFC 5322 allowed character set.
func isValidLocalPartChars(local string) bool {
	for _, c := range local {
		if !isValidLocalChar(c) {
			return false
		}
	}

	return true
}

// isValidLocalChar checks if a character is valid in local part.
// RFC 5322 section 3.2.3 defines atext (atom text) as:
//
//	ALPHA / DIGIT / "!" / "#" / "$" / "%" / "&" / "'" / "*" /
//	"+" / "-" / "/" / "=" / "?" / "^" / "_" / "`" / "{" /
//	"|" / "}" / "~" / "."
func isValidLocalChar(c rune) bool {
	// Check alphanumeric characters first (most common case)
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
		return true
	}

	return isValidLocalSpecialChar(c)
}

// isValidLocalSpecialChar checks if a character is a valid special character in local part.
// These are the special characters allowed by RFC 5322 in addition to alphanumeric.
//
// Special characters and their ASCII codes:
//
//	! (33)  # (35)  $ (36)  % (37)  & (38)  ' (39)  * (42)  + (43)
//	- (45)  . (46)  / (47)  = (61)  ? (63)  ^ (94)  _ (95)  ` (96)
//	{ (123) | (124) } (125) ~ (126)
func isValidLocalSpecialChar(c rune) bool {
	// Using switch for readability and performance
	switch c {
	case '.', '_', '-', '+', '=', '!', '#', '$', '%', '&', '\'', '*', '/', '?', '^', '`', '{', '|', '}', '~':
		return true
	default:
		return false
	}
}

// isValidDomainPart validates the domain part (after @) of an email.
// Domain validation follows RFC 1035 (Domain Names) and RFC 5321 (SMTP).
//
// Domain format: label.label[.label...]
//
// Requirements:
//   - Total length: 1-253 characters (RFC 1035 section 2.3.4)
//   - Must contain at least one dot
//   - Must have at least 2 labels (e.g., "example.com")
//   - Cannot start or end with dot or hyphen
//
// Valid examples:
//
//	"example.com"         -> 2 labels
//	"mail.example.com"    -> 3 labels
//	"a.b.c.d.example.com" -> 6 labels
//
// Invalid examples:
//
//	"localhost"           -> no dot
//	".example.com"        -> starts with dot
//	"example.com."        -> ends with dot
//	"-example.com"        -> starts with hyphen
func isValidDomainPart(domain string) bool {
	// Length check: 1-253 characters (RFC 1035 section 2.3.4)
	if domain == "" || len(domain) > 253 {
		return false
	}

	// Must contain at least one dot (to have multiple labels)
	// This ensures we have a proper domain like "example.com" not just "localhost"
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
	// These are invalid at domain boundaries
	if domain[0] == '.' || domain[len(domain)-1] == '.' ||
		domain[0] == '-' || domain[len(domain)-1] == '-' {
		return false
	}

	return validateDomainLabels(domain)
}

// validateDomainLabels parses and validates each domain label.
// A label is a segment between dots in the domain.
//
// Example parsing "mail.example.com":
//
//	Position: 0123456789012345
//	Labels:   mail|example|com
//	Dots at:      4       11
//
// Each label is validated for:
//   - Non-empty
//   - Valid length (1-63 characters)
//   - Valid characters and format
func validateDomainLabels(domain string) bool {
	labelCount := 0
	start := 0

	// Parse domain into labels by iterating character by character
	for i := 0; i <= len(domain); i++ {
		// Process label when we hit a dot or reach the end
		if i != len(domain) && domain[i] != '.' {
			continue
		}

		if i == start {
			// Empty label (consecutive dots or dot at start/end)
			return false
		}

		// Extract label from start position to current position
		label := domain[start:i]
		labelCount++

		// Validate this label
		if !isValidDomainLabel(label) {
			return false
		}

		// Move start to character after the dot
		start = i + 1
	}

	// Must have at least 2 labels to form a valid domain
	// e.g., "example.com" has 2 labels, "localhost" has only 1
	return labelCount >= 2
}

// isValidDomainLabel validates a single domain label.
// RFC 1035 section 2.3.1 defines label rules:
//   - Maximum 63 octets
//   - Must not start or end with hyphen
//   - Contains only letters, digits, and hyphens
//
// Valid examples:
//
//	"example"    -> alphanumeric only
//	"ex-ample"   -> with hyphen in middle
//	"example123" -> alphanumeric mix
//
// Invalid examples:
//
//	"-example"   -> starts with hyphen
//	"example-"   -> ends with hyphen
//	"ex ample"   -> contains space
//	"ex_ample"   -> contains underscore (not allowed in domain)
func isValidDomainLabel(label string) bool {
	// Length check: 1-63 characters (RFC 1035 section 2.3.1)
	if label == "" || len(label) > 63 {
		return false
	}

	// RFC 952: Labels must not start or end with hyphen
	if label[0] == '-' || label[len(label)-1] == '-' {
		return false
	}

	return isValidDomainLabelChars(label)
}

// isValidDomainLabelChars checks if all characters in label are valid.
// Each character must be alphanumeric or hyphen.
func isValidDomainLabelChars(label string) bool {
	for _, c := range label {
		if !isValidDomainChar(c) {
			return false
		}
	}

	return true
}

// isValidDomainChar checks if a character is valid in domain label.
// RFC 1035: domain labels consist of letters, digits, and hyphens.
// Note: Underscore is NOT allowed in domain names (unlike local part).
//
// Character ranges:
//
//	a-z: 97-122
//	A-Z: 65-90
//	0-9: 48-57
//	-:   45
func isValidDomainChar(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') ||
		(c >= '0' && c <= '9') || c == '-'
}
