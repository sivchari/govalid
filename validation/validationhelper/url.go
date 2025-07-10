package validationhelper

// validSchemes contains all supported URL schemes.
var validSchemes = map[string]bool{
	"http":   true,
	"https":  true,
	"ftp":    true,
	"ftps":   true,
	"ssh":    true,
	"sftp":   true,
	"smtp":   true,
	"smtps":  true,
	"imap":   true,
	"imaps":  true,
	"pop3":   true,
	"pop3s":  true,
	"telnet": true,
	"file":   true,
	"data":   true,
	"ws":     true,
	"wss":    true,
	"git":    true,
	"svn":    true,
	"ldap":   true,
	"ldaps":  true,
	"mailto": true,
	"news":   true,
	"nntp":   true,
	"irc":    true,
	"ircs":   true,
	"rtsp":   true,
	"rtmp":   true,
	"sip":    true,
	"sips":   true,
	"xmpp":   true,
}

// schemesNotRequiringHost contains schemes that don't require a host part.
var schemesNotRequiringHost = map[string]bool{
	"mailto": true,
	"news":   true,
	"nntp":   true,
	"data":   true,
	"file":   true, // file:// can be file:/path or file:///path
}

// IsValidURL validates if a string is a valid URL format.
// This implementation prioritizes correctness while maintaining high performance.
func IsValidURL(input string) bool {
	// Empty string is obviously invalid
	if input == "" {
		return false
	}

	// Find the colon that separates scheme from the rest
	colonPos := findSchemeEnd(input)
	if colonPos == -1 || colonPos == 0 {
		return false
	}

	// Extract and validate scheme
	scheme := input[:colonPos]
	if !validSchemes[scheme] {
		return false
	}

	// Check for spaces and control characters (invalid in URLs)
	if hasInvalidChars(input) {
		return false
	}

	// Handle schemes that don't require host
	if schemesNotRequiringHost[scheme] {
		return validateSchemeWithoutHost(input, colonPos)
	}

	// Handle schemes that require host (scheme://host)
	return validateSchemeWithHost(input, colonPos)
}

// findSchemeEnd finds the position of the colon that ends the scheme.
// Returns -1 if no valid scheme colon is found.
func findSchemeEnd(input string) int {
	inputLen := len(input)

	// RFC 3986: scheme = ALPHA *( ALPHA / DIGIT / "+" / "-" / "." )
	// We start from position 1 since scheme must start with ALPHA
	for i := 1; i < inputLen; i++ {
		char := input[i]
		if char == ':' {
			return i
		}

		if !isValidSchemeChar(char) {
			return -1
		}
	}

	return -1
}

// isValidSchemeChar checks if a character is valid in a URL scheme.
func isValidSchemeChar(char byte) bool {
	return (char >= 'a' && char <= 'z') ||
		(char >= 'A' && char <= 'Z') ||
		(char >= '0' && char <= '9') ||
		char == '+' || char == '-' || char == '.'
}

// hasInvalidChars checks if the input contains spaces or control characters.
func hasInvalidChars(input string) bool {
	for i := 0; i < len(input); i++ {
		char := input[i]
		// Check for spaces
		if char == ' ' {
			return true
		}
		// Check for control characters (0-31 and 127)
		if char < 32 || char == 127 {
			return true
		}
	}

	return false
}

// validateSchemeWithoutHost validates URLs for schemes that don't require a host.
//
// Expected format: scheme:opaque-part
//
// Example breakdown:
//
//	"mailto:user@example.com"
//	Position: 0123456789...
//	         mailto:user@example.com
//	         ↑     ↑
//	         0     6 7
//	colonPos = 6 (position of ":")
//	colonPos+1 = 7 (opaque part start)
//
// Validation steps:
//  1. Ensure there's content after the colon
//  2. Accept any content (simplified validation)
func validateSchemeWithoutHost(input string, colonPos int) bool {
	// Step 1: Must have something after the colon
	// Examples:
	//   "mailto:"              → colonPos=6, +1=7, len=7 → 7>=7 → false (empty)
	//   "mailto:x"             → colonPos=6, +1=7, len=8 → 7>=8 → true (valid)
	//   "data:text/plain,Hi"   → colonPos=4, +1=5, len=16 → 5>=16 → true (valid)
	if colonPos+1 >= len(input) {
		return false
	}

	return true
}

// validateSchemeWithHost validates URLs for schemes that require a host.
//
// Expected format: scheme://host[/path][?query][#fragment]
//
// Example breakdown:
//
//	"http://example.com"
//	Position: 0123456789...
//	         http://example.com
//	         ↑   ↑ ↑ ↑
//	         0   4 5 6 7
//	colonPos = 4 (position of ":")
//	colonPos+1 = 5 (first "/")
//	colonPos+2 = 6 (second "/")
//	colonPos+3 = 7 (host start position)
//
// Validation steps:
//  1. Check if "://" pattern exists and has room for host
//  2. Verify actual "://" characters
//  3. Ensure host exists after "://"
//  4. Validate host start character
func validateSchemeWithHost(input string, colonPos int) bool {
	inputLen := len(input)

	// Step 1: Check if we have enough characters for "://" + at least 1 host char
	// Examples:
	//   "http:"     → colonPos=4, +3=7, len=5 → 7>=5 → false (too short)
	//   "http://"   → colonPos=4, +3=7, len=7 → 7>=7 → false (no host)
	//   "http://x"  → colonPos=4, +3=7, len=8 → 7>=8 → true (continue)
	if colonPos+3 >= inputLen {
		return false
	}

	if input[colonPos+1] != '/' || input[colonPos+2] != '/' {
		return false
	}

	hostStart := colonPos + 3
	if hostStart >= inputLen {
		return false
	}

	return isValidHostStart(input[hostStart])
}

// isValidHostStart checks if the first character of a host is valid.
func isValidHostStart(char byte) bool {
	return (char >= 'a' && char <= 'z') ||
		(char >= 'A' && char <= 'Z') ||
		(char >= '0' && char <= '9') ||
		char == '[' // IPv6 addresses start with [
}
