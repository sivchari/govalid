// Package validationhelper provides validation helper functions for govalid.
package validationhelper

// IsValidEmail validates email format manually for maximum performance.
// Validates basic email structure: local@domain format.
func IsValidEmail(email string) bool {
	// Basic length check
	if len(email) < 5 || len(email) > 254 { // a@b.c = 5 chars minimum, RFC 5321 limit
		return false
	}
	
	// Must contain exactly one @
	atIndex := -1
	atCount := 0
	for i, c := range email {
		if c == '@' {
			atIndex = i
			atCount++
		}
	}
	if atCount != 1 || atIndex <= 0 || atIndex >= len(email)-1 {
		return false
	}
	
	// Validate local part (before @)
	local := email[:atIndex]
	if len(local) == 0 || len(local) > 64 { // RFC 5321 limit
		return false
	}
	
	// Check for consecutive dots or leading/trailing dots
	if local[0] == '.' || local[len(local)-1] == '.' {
		return false
	}
	for i := 0; i < len(local)-1; i++ {
		if local[i] == '.' && local[i+1] == '.' {
			return false
		}
	}
	
	// Check allowed characters in local part
	for _, c := range local {
		if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || 
			 (c >= '0' && c <= '9') || c == '.' || c == '_' || c == '-' || 
			 c == '+' || c == '=' || c == '!' || c == '#' || c == '$' || 
			 c == '%' || c == '&' || c == '\'' || c == '*' || c == '/' || 
			 c == '?' || c == '^' || c == 96 || c == '{' || c == '|' || 
			 c == '}' || c == '~') { // c == 96 is backtick
			return false
		}
	}
	
	// Validate domain part (after @)
	domain := email[atIndex+1:]
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
	
	// Parse domain labels manually to avoid allocations
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
			
			// Validate label length
			if len(label) > 63 { // RFC 1035 limit
				return false
			}
			
			// Cannot start or end with hyphen
			if label[0] == '-' || label[len(label)-1] == '-' {
				return false
			}
			
			// Check allowed characters (alphanumeric + hyphen)
			for _, c := range label {
				if !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || 
					 (c >= '0' && c <= '9') || c == '-') {
					return false
				}
			}
			
			start = i + 1
		}
	}
	
	// Must have at least 2 labels (e.g., "example.com")
	if labelCount < 2 {
		return false
	}
	
	return true
}