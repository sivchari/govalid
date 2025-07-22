package validationhelper

import "testing"

func TestIsValidURL(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		// Valid URLs
		{"valid_https", "https://example.com", true},
		{"valid_http", "http://example.com", true},
		{"valid_ftp", "ftp://files.example.com", true},
		{"valid_ws", "ws://socket.example.com", true},
		{"valid_mailto", "mailto:user@example.com", true},
		{"valid_data", "data:text/plain,Hello", true},
		{"valid_file", "file:/path/to/file", true},
		{"valid_with_port", "http://example.com:8080", true},
		{"valid_with_path", "https://example.com/path", true},
		{"valid_with_query", "https://example.com?q=test", true},
		{"valid_ipv6", "http://[::1]", true},
		{"valid_ipv4", "http://192.168.1.1", true},

		// Invalid URLs
		{"empty_string", "", false},
		{"no_scheme", "example.com", false},
		{"no_colon", "httpexample.com", false},
		{"invalid_scheme_char", "ht@tp://example.com", false},
		{"unknown_scheme", "unknown://example.com", false},
		{"has_space", "http://example .com", false},
		{"schemes_with_host_no_slashes", "http:", false},
		{"schemes_with_host_one_slash", "http:/", false},
		{"schemes_with_host_no_host", "http://", false},
		{"schemes_without_host_empty", "mailto:", false},
		{"schemes_without_host_empty_data", "data:", false},
		{"invalid_host_start_dot", "http://.example.com", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidURL(tt.input)
			if result != tt.expected {
				t.Errorf("IsValidURL(%q) = %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestFindSchemeEnd(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"valid_http", "http://example.com", 4},
		{"valid_https", "https://example.com", 5},
		{"valid_ftp", "ftp://files.example.com", 3},
		{"valid_mailto", "mailto:user@example.com", 6},
		{"valid_data", "data:text/plain,Hello", 4},
		{"valid_ws", "ws://socket.example.com", 2},
		{"valid_scheme_with_plus", "svn+ssh://example.com", 7},
		{"valid_scheme_with_dash", "svn-ssh://example.com", 7},
		{"valid_scheme_with_dot", "svn.ssh://example.com", 7},

		{"no_colon", "httpexample.com", -1},
		{"empty_string", "", -1},
		{"single_char", "h", -1},
		{"colon_at_start", ":example.com", -1},
		{"invalid_scheme_char", "ht@tp://example.com", -1},
		{"scheme_with_number", "http2://example.com", 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findSchemeEnd(tt.input)
			if result != tt.expected {
				t.Errorf("findSchemeEnd(%q) = %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestIsValidSchemeChar(t *testing.T) {
	tests := []struct {
		name     string
		char     byte
		expected bool
	}{
		// Valid characters
		{"lowercase_a", 'a', true},
		{"lowercase_z", 'z', true},
		{"uppercase_A", 'A', true},
		{"uppercase_Z", 'Z', true},
		{"digit_0", '0', true},
		{"digit_9", '9', true},
		{"plus", '+', true},
		{"dash", '-', true},
		{"dot", '.', true},

		// Invalid characters
		{"at_sign", '@', false},
		{"space", ' ', false},
		{"colon", ':', false},
		{"slash", '/', false},
		{"question", '?', false},
		{"hash", '#', false},
		{"underscore", '_', false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidSchemeChar(tt.char)
			if result != tt.expected {
				t.Errorf("isValidSchemeChar(%q) = %v, expected %v", tt.char, result, tt.expected)
			}
		})
	}
}

func TestHasInvalidChars(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"no_invalid_chars", "https://example.com", false},
		{"space_in_middle", "https://example .com", true},
		{"space_at_start", " https://example.com", true},
		{"space_at_end", "https://example.com ", true},
		{"multiple_spaces", "https://example .com /path", true},
		{"empty_string", "", false},
		{"only_space", " ", true},
		{"tab_character", "https://example\t.com", true},     // tab is control character
		{"newline_character", "https://example\n.com", true}, // newline is control character
		{"null_character", "https://example\x00.com", true},  // null is control character
		{"del_character", "https://example\x7f.com", true},   // DEL is control character
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := hasInvalidChars(tt.input)
			if result != tt.expected {
				t.Errorf("hasInvalidChars(%q) = %v, expected %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestValidateSchemeWithoutHost(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		colonPos int
		expected bool
	}{
		{"valid_mailto", "mailto:user@example.com", 6, true},
		{"valid_data", "data:text/plain,Hello", 4, true},
		{"valid_news", "news:comp.lang.go", 4, true},
		{"valid_file", "file:/path/to/file", 4, true},
		{"valid_single_char", "mailto:x", 6, true},

		{"empty_after_colon", "mailto:", 6, false},
		{"empty_after_colon_data", "data:", 4, false},
		{"colon_at_end", "mailto", 5, false}, // colonPos beyond string
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validateSchemeWithoutHost(tt.input, tt.colonPos)
			if result != tt.expected {
				t.Errorf("validateSchemeWithoutHost(%q, %d) = %v, expected %v",
					tt.input, tt.colonPos, result, tt.expected)
			}
		})
	}
}

func TestValidateSchemeWithHost(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		colonPos int
		expected bool
	}{
		{"valid_http", "http://example.com", 4, true},
		{"valid_https", "https://example.com", 5, true},
		{"valid_ftp", "ftp://files.example.com", 3, true},
		{"valid_ws", "ws://socket.example.com", 2, true},
		{"valid_single_char_host", "http://x", 4, true},
		{"valid_ipv6", "http://[::1]", 4, true},
		{"valid_ipv4", "http://192.168.1.1", 4, true},

		{"no_slashes", "http:", 4, false},
		{"one_slash", "http:/", 4, false},
		{"no_host", "http://", 4, false},
		{"invalid_host_start_dot", "http://.example.com", 4, false},
		{"too_short_for_host", "http://", 4, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validateSchemeWithHost(tt.input, tt.colonPos)
			if result != tt.expected {
				t.Errorf("validateSchemeWithHost(%q, %d) = %v, expected %v",
					tt.input, tt.colonPos, result, tt.expected)
			}
		})
	}
}

func TestIsValidHostStart(t *testing.T) {
	tests := []struct {
		name     string
		char     byte
		expected bool
	}{
		// Valid characters
		{"lowercase_a", 'a', true},
		{"lowercase_z", 'z', true},
		{"uppercase_A", 'A', true},
		{"uppercase_Z", 'Z', true},
		{"digit_0", '0', true},
		{"digit_9", '9', true},
		{"ipv6_bracket", '[', true},

		// Invalid characters
		{"dot", '.', false},
		{"dash", '-', false},
		{"space", ' ', false},
		{"slash", '/', false},
		{"colon", ':', false},
		{"at_sign", '@', false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidHostStart(tt.char)
			if result != tt.expected {
				t.Errorf("isValidHostStart(%q) = %v, expected %v", tt.char, result, tt.expected)
			}
		})
	}
}

// FuzzIsValidURL performs fuzz testing on URL validation.
func FuzzIsValidURL(f *testing.F) {
	addURLFuzzSeeds(f)

	f.Fuzz(func(t *testing.T, url string) {
		// The function should never panic, regardless of input
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("IsValidURL panicked on input %q: %v", url, r)
			}
		}()

		result := IsValidURL(url)

		// Basic invariants that should always hold
		if result {
			validateURLStructure(t, url)
		}

		// Test that the function is deterministic
		result2 := IsValidURL(url)
		if result != result2 {
			t.Errorf("IsValidURL(%q) is not deterministic: got %v then %v", url, result, result2)
		}
	})
}

func addURLFuzzSeeds(f *testing.F) {
	f.Helper()

	seeds := []string{
		"https://example.com",
		"http://example.com",
		"ftp://files.example.com",
		"ws://socket.example.com",
		"mailto:user@example.com",
		"data:text/plain,Hello",
		"file:/path/to/file",
		"http://example.com:8080",
		"https://example.com/path",
		"https://example.com?q=test",
		"http://[::1]",
		"http://192.168.1.1",
		"https://user:pass@example.com",
		"https://example.com/path/to/resource?query=value#fragment",
		"",
		"example.com",
		"httpexample.com",
		"ht@tp://example.com",
		"unknown://example.com",
		"http://example .com",
		"http:",
		"http:/",
		"http://",
		"mailto:",
		"data:",
		"http://.example.com",
		"http://example.com.",
		"http://example.com:abc",
		"http://example.com:-1",
		"http://example.com:999999",
		"http://example..com",
		"https://",
		"https://.",
		"https://-example.com",
		"https://example-.com",
		"https://ex ample.com",
		"https://example.com/path with spaces",
		"https://example.com\npath",
		"https://example.com\tpath",
		"https://example.com/path\x00",
		"https://[invalid-ipv6]",
		"https://[::1::2]",
		"https://999.999.999.999",
		"https://256.1.1.1",
		"file://",
		"file:///",
		"file://localhost/path",
		"ldap://example.com",
		"custom://test",
		"javascript:alert('xss')",
		"vbscript:msgbox('xss')",
	}

	for _, seed := range seeds {
		f.Add(seed)
	}
}

func validateURLStructure(t *testing.T, url string) {
	t.Helper()

	// Valid URLs must contain a colon for the scheme separator
	if !validateURLHasColon(t, url) {
		return
	}

	validateURLScheme(t, url)
	validateURLCharacters(t, url)
}

func validateURLHasColon(t *testing.T, url string) bool {
	t.Helper()

	hasColon := false

	for _, c := range url {
		if c == ':' {
			hasColon = true

			break
		}
	}

	if !hasColon {
		t.Errorf("IsValidURL(%q) returned true but contains no colon", url)

		return false
	}

	return true
}

func validateURLScheme(t *testing.T, url string) {
	t.Helper()

	colonPos := -1

	for i, c := range url {
		if c == ':' {
			colonPos = i

			break
		}
	}

	if colonPos > 0 {
		scheme := url[:colonPos]
		if scheme == "" {
			t.Errorf("IsValidURL(%q) returned true but has empty scheme", url)

			return
		}
		// First character must be a letter
		if scheme != "" {
			first := scheme[0]
			if (first >= 'a' && first <= 'z') || (first >= 'A' && first <= 'Z') {
				return
			}

			t.Errorf("IsValidURL(%q) returned true but scheme starts with invalid char %c", url, first)
		}
	}
}

func validateURLCharacters(t *testing.T, url string) {
	t.Helper()

	// URLs should not contain control characters or spaces
	for i, c := range url {
		if c < 32 || c == 127 {
			t.Errorf("IsValidURL(%q) returned true but contains control character %d at position %d", url, c, i)
		}

		if c == ' ' {
			t.Errorf("IsValidURL(%q) returned true but contains space at position %d", url, i)
		}
	}
}
