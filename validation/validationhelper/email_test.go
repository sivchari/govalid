package validationhelper

import (
	"strings"
	"testing"
)

func TestIsValidEmail(t *testing.T) {
	t.Run("valid_emails", testValidEmails)
	t.Run("invalid_emails", testInvalidEmails)
	t.Run("length_boundary_tests", testEmailLengthBoundaries)
}

func testValidEmails(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		expected bool
	}{
		{"simple_email", "user@example.com", true},
		{"with_subdomain", "user@mail.example.com", true},
		{"with_numbers", "user123@example456.com", true},
		{"with_dots", "user.name@example.com", true},
		{"with_plus", "user+tag@example.com", true},
		{"with_hyphen", "user-name@ex-ample.com", true},
		{"with_underscore", "user_name@example.com", true},
		{"mixed_case", "UserName@Example.Com", true},
		{"minimal_valid", "a@b.c", true},
		{"long_domain", "user@sub1.sub2.sub3.example.com", true},
		{"special_chars_local", "user!#$%&'*+-/=?^_`{|}~@example.com", true},
		{"numeric_local", "123456@example.com", true},
		{"numeric_domain", "user@123.456.com", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidEmail(tt.email)
			if result != tt.expected {
				t.Errorf("IsValidEmail(%q) = %v, expected %v", tt.email, result, tt.expected)
			}
		})
	}
}

func testInvalidEmails(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		expected bool
	}{
		// Invalid emails - structure
		{"empty_string", "", false},
		{"no_at_symbol", "userexample.com", false},
		{"no_domain", "user@", false},
		{"no_local", "@example.com", false},
		{"multiple_at", "user@@example.com", false},
		{"at_in_middle", "user@middle@example.com", false},
		{"no_domain_dot", "user@localhost", false},
		{"spaces_in_email", "user name@example.com", false},
		{"trailing_dot_domain", "user@example.com.", false},
		{"leading_dot_domain", "user@.example.com", false},
		{"consecutive_dots_domain", "user@example..com", false},

		// Invalid emails - local part
		{"leading_dot_local", ".user@example.com", false},
		{"trailing_dot_local", "user.@example.com", false},
		{"consecutive_dots_local", "user..name@example.com", false},
		{"invalid_char_local", "user<name>@example.com", false},
		{"empty_local", "@example.com", false},
		{"too_long_local", "a123456789012345678901234567890123456789012345678901234567890123456789@example.com", false},

		// Invalid emails - domain part
		{"domain_starts_hyphen", "user@-example.com", false},
		{"domain_ends_hyphen", "user@example-.com", false},
		{"label_starts_hyphen", "user@ex-ample.-com", false},
		{"label_ends_hyphen", "user@ex-ample.com-", false},
		{"underscore_in_domain", "user@ex_ample.com", false},
		{"too_long_label", "user@" + generateString(64) + ".com", false},
		{"empty_label", "user@example..com", false},
		{"single_label", "user@example", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidEmail(tt.email)
			if result != tt.expected {
				t.Errorf("IsValidEmail(%q) = %v, expected %v", tt.email, result, tt.expected)
			}
		})
	}
}

func testEmailLengthBoundaries(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		expected bool
	}{
		{"max_valid_length", "a@" + generateValidDomain(246) + ".com", true}, // 254 total
		{"too_long_email", "a@" + generateValidDomain(249) + ".com", false},  // 255 total
		{"max_local_length", generateString(64) + "@example.com", true},
		{"too_long_local", generateString(65) + "@example.com", false},
		{"max_domain_length", "a@" + generateValidDomain(247) + ".com", true}, // 253 domain
		{"too_long_domain", "a@" + generateValidDomain(250) + ".com", false},  // 254 domain
		{"max_label_length", "user@" + generateString(63) + ".com", true},
		{"too_long_label", "user@" + generateString(64) + ".com", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidEmail(tt.email)
			if result != tt.expected {
				t.Errorf("IsValidEmail(%q) = %v, expected %v", tt.email, result, tt.expected)
			}
		})
	}
}

func TestFindAtSymbol(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		expected int
	}{
		{"normal_email", "user@example.com", 4},
		{"early_at", "u@example.com", 1},
		{"late_at", "username@e.com", 8},
		{"no_at", "userexample.com", -1},
		{"multiple_at", "user@@example.com", -1},
		{"at_at_start", "@example.com", -1},
		{"at_at_end", "user@", -1},
		{"only_at", "@", -1},
		{"two_separate_at", "user@middle@example.com", -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findAtSymbol(tt.email)
			if result != tt.expected {
				t.Errorf("findAtSymbol(%q) = %v, expected %v", tt.email, result, tt.expected)
			}
		})
	}
}

func TestIsValidLocalPart(t *testing.T) {
	tests := []struct {
		name     string
		local    string
		expected bool
	}{
		// Valid local parts
		{"simple", "user", true},
		{"with_dot", "user.name", true},
		{"with_numbers", "user123", true},
		{"all_numbers", "123456", true},
		{"with_plus", "user+tag", true},
		{"with_hyphen", "user-name", true},
		{"with_underscore", "user_name", true},
		{"special_chars", "user!#$%&'*+-/=?^_`{|}~", true},
		{"max_length", generateString(64), true},
		{"single_char", "a", true},

		// Invalid local parts
		{"empty", "", false},
		{"too_long", generateString(65), false},
		{"leading_dot", ".user", false},
		{"trailing_dot", "user.", false},
		{"consecutive_dots", "user..name", false},
		{"invalid_char_angle", "user<name>", false},
		{"invalid_char_paren", "user(name)", false},
		{"invalid_char_comma", "user,name", false},
		{"invalid_char_colon", "user:name", false},
		{"invalid_char_semicolon", "user;name", false},
		{"invalid_char_bracket", "user[name]", false},
		{"invalid_char_backslash", "user\\name", false},
		{"invalid_char_quote", "user\"name", false},
		{"invalid_char_space", "user name", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidLocalPart(tt.local)
			if result != tt.expected {
				t.Errorf("isValidLocalPart(%q) = %v, expected %v", tt.local, result, tt.expected)
			}
		})
	}
}

func TestIsValidLocalPartFormat(t *testing.T) {
	tests := []struct {
		name     string
		local    string
		expected bool
	}{
		{"no_dots", "username", true},
		{"single_dot", "user.name", true},
		{"multiple_dots", "user.middle.name", true},
		{"leading_dot", ".username", false},
		{"trailing_dot", "username.", false},
		{"consecutive_dots", "user..name", false},
		{"dot_at_both_ends", ".username.", false},
		{"only_dots", "...", false},
		{"single_char", "a", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidLocalPartFormat(tt.local)
			if result != tt.expected {
				t.Errorf("isValidLocalPartFormat(%q) = %v, expected %v", tt.local, result, tt.expected)
			}
		})
	}
}

func TestIsValidLocalChar(t *testing.T) {
	tests := []struct {
		name     string
		char     rune
		expected bool
	}{
		// Valid characters
		{"lowercase_a", 'a', true},
		{"lowercase_z", 'z', true},
		{"uppercase_A", 'A', true},
		{"uppercase_Z", 'Z', true},
		{"digit_0", '0', true},
		{"digit_9", '9', true},
		{"dot", '.', true},
		{"underscore", '_', true},
		{"hyphen", '-', true},
		{"plus", '+', true},
		{"equals", '=', true},
		{"exclamation", '!', true},
		{"hash", '#', true},
		{"dollar", '$', true},
		{"percent", '%', true},
		{"ampersand", '&', true},
		{"apostrophe", '\'', true},
		{"asterisk", '*', true},
		{"slash", '/', true},
		{"question", '?', true},
		{"caret", '^', true},
		{"backtick", '`', true},
		{"left_brace", '{', true},
		{"pipe", '|', true},
		{"right_brace", '}', true},
		{"tilde", '~', true},

		// Invalid characters
		{"space", ' ', false},
		{"at", '@', false},
		{"left_paren", '(', false},
		{"right_paren", ')', false},
		{"comma", ',', false},
		{"colon", ':', false},
		{"semicolon", ';', false},
		{"less_than", '<', false},
		{"greater_than", '>', false},
		{"left_bracket", '[', false},
		{"right_bracket", ']', false},
		{"backslash", '\\', false},
		{"quote", '"', false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidLocalChar(tt.char)
			if result != tt.expected {
				t.Errorf("isValidLocalChar(%q) = %v, expected %v", tt.char, result, tt.expected)
			}
		})
	}
}

func TestIsValidDomainPart(t *testing.T) {
	tests := []struct {
		name     string
		domain   string
		expected bool
	}{
		// Valid domains
		{"simple", "example.com", true},
		{"with_subdomain", "mail.example.com", true},
		{"multiple_subdomains", "a.b.c.example.com", true},
		{"with_numbers", "example123.com", true},
		{"with_hyphen", "ex-ample.com", true},
		{"hyphen_in_label", "mail-server.example.com", true},
		{"mixed_case", "Example.Com", true},
		{"two_char_tld", "example.io", true},
		{"long_tld", "example.technology", true},
		{"numeric_labels", "123.456.com", true},
		{"max_length", generateValidDomain(241) + ".example.com", true}, // 253 total

		// Invalid domains
		{"empty", "", false},
		{"too_long", generateValidDomain(242) + ".example.com", false}, // 254 total
		{"no_dot", "localhost", false},
		{"single_label", "com", false},
		{"leading_dot", ".example.com", false},
		{"trailing_dot", "example.com.", false},
		{"leading_hyphen", "-example.com", false},
		{"trailing_hyphen", "example-.com", false},
		{"consecutive_dots", "example..com", false},
		{"empty_label", "example..com", false},
		{"label_starts_hyphen", "ex.-ample.com", false},
		{"label_ends_hyphen", "ex.ample-.com", false},
		{"underscore", "ex_ample.com", false},
		{"special_char", "ex@mple.com", false},
		{"space", "ex ample.com", false},
		{"too_long_label", generateString(64) + ".com", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidDomainPart(tt.domain)
			if result != tt.expected {
				t.Errorf("isValidDomainPart(%q) = %v, expected %v", tt.domain, result, tt.expected)
			}
		})
	}
}

func TestValidateDomainLabels(t *testing.T) {
	tests := []struct {
		name     string
		domain   string
		expected bool
	}{
		{"two_labels", "example.com", true},
		{"three_labels", "mail.example.com", true},
		{"many_labels", "a.b.c.d.e.f.com", true},
		{"single_label", "localhost", false},
		{"empty_label", "example..com", false},
		{"consecutive_dots", "example...com", false},
		{"label_too_long", "a" + generateString(64) + ".com", false},
		{"valid_max_label", generateString(63) + ".com", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := validateDomainLabels(tt.domain)
			if result != tt.expected {
				t.Errorf("validateDomainLabels(%q) = %v, expected %v", tt.domain, result, tt.expected)
			}
		})
	}
}

func TestIsValidDomainLabel(t *testing.T) {
	tests := []struct {
		name     string
		label    string
		expected bool
	}{
		{"simple", "example", true},
		{"with_numbers", "example123", true},
		{"all_numbers", "123456", true},
		{"with_hyphen", "ex-ample", true},
		{"multiple_hyphens", "ex-am-ple", true},
		{"max_length", generateString(63), true},
		{"single_char", "a", true},

		{"too_long", generateString(64), false},
		{"starts_hyphen", "-example", false},
		{"ends_hyphen", "example-", false},
		{"underscore", "ex_ample", false},
		{"special_char", "ex@mple", false},
		{"space", "ex ample", false},
		{"dot", "ex.ample", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidDomainLabel(tt.label)
			if result != tt.expected {
				t.Errorf("isValidDomainLabel(%q) = %v, expected %v", tt.label, result, tt.expected)
			}
		})
	}
}

func TestIsValidDomainChar(t *testing.T) {
	tests := []struct {
		name     string
		char     rune
		expected bool
	}{
		// Valid characters
		{"lowercase_a", 'a', true},
		{"lowercase_z", 'z', true},
		{"uppercase_A", 'A', true},
		{"uppercase_Z", 'Z', true},
		{"digit_0", '0', true},
		{"digit_9", '9', true},
		{"hyphen", '-', true},

		// Invalid characters
		{"underscore", '_', false},
		{"dot", '.', false},
		{"space", ' ', false},
		{"at", '@', false},
		{"special_char", '!', false},
		{"plus", '+', false},
		{"equals", '=', false},
		{"slash", '/', false},
		{"backslash", '\\', false},
		{"quote", '"', false},
		{"apostrophe", '\'', false},
		{"left_paren", '(', false},
		{"right_paren", ')', false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidDomainChar(tt.char)
			if result != tt.expected {
				t.Errorf("isValidDomainChar(%q) = %v, expected %v", tt.char, result, tt.expected)
			}
		})
	}
}

// FuzzIsValidEmail performs fuzz testing on email validation.
func FuzzIsValidEmail(f *testing.F) {
	addEmailFuzzSeeds(f)

	f.Fuzz(func(t *testing.T, email string) {
		// The function should never panic, regardless of input
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("IsValidEmail panicked on input %q: %v", email, r)
			}
		}()

		result := IsValidEmail(email)

		// Basic invariants that should always hold
		if result {
			validateEmailStructure(t, email)
		}

		// Test that the function is deterministic
		result2 := IsValidEmail(email)
		if result != result2 {
			t.Errorf("IsValidEmail(%q) is not deterministic: got %v then %v", email, result, result2)
		}
	})
}

func addEmailFuzzSeeds(f *testing.F) {
	f.Helper()

	seeds := []string{
		"user@example.com",
		"test.email@domain.org",
		"user+tag@mail.example.com",
		"invalid@",
		"@invalid.com",
		"user@@example.com",
		"user@domain..com",
		"user@.domain.com",
		"user@domain.com.",
		"user..name@example.com",
		".user@example.com",
		"user.@example.com",
		"user@localhost",
		"user@domain",
		"user@example.c",
		"user@192.168.1.1",
		"user@[192.168.1.1]",
		"user name@example.com",
		"user@ex ample.com",
		"",
		"a@b.c",
		generateString(300) + "@example.com",
		"user@" + generateString(300) + ".com",
		"user@example." + generateString(300),
	}

	for _, seed := range seeds {
		f.Add(seed)
	}
}

func validateEmailStructure(t *testing.T, email string) {
	t.Helper()

	// Valid emails must have basic structural requirements
	if len(email) < 5 || len(email) > 254 {
		t.Errorf("IsValidEmail(%q) returned true but length is invalid: %d", email, len(email))
	}

	// Must contain exactly one @ symbol
	atCount := 0

	for _, c := range email {
		if c == '@' {
			atCount++
		}
	}

	if atCount != 1 {
		t.Errorf("IsValidEmail(%q) returned true but @ count is %d", email, atCount)
	}

	// Must have at least one dot in domain part
	validateEmailDomainDot(t, email)
}

func validateEmailDomainDot(t *testing.T, email string) {
	t.Helper()

	atIndex := -1

	for i, c := range email {
		if c == '@' {
			atIndex = i

			break
		}
	}

	if atIndex > 0 {
		domain := email[atIndex+1:]
		hasDot := false

		for _, c := range domain {
			if c == '.' {
				hasDot = true

				break
			}
		}

		if !hasDot {
			t.Errorf("IsValidEmail(%q) returned true but domain has no dot", email)
		}
	}
}

// Helper function to generate strings of specific length.
func generateString(length int) string {
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = byte('a' + (i % 26))
	}

	return string(result)
}

// Helper function to generate valid domain names of specific length.
// Creates domains with labels that respect the 63 character limit.
func generateValidDomain(length int) string {
	if length <= 0 {
		return ""
	}

	var result strings.Builder

	remaining := length
	labelCount := 0

	for remaining > 0 {
		if labelCount > 0 {
			result.WriteByte('.')

			remaining--
			if remaining <= 0 {
				break
			}
		}

		// Each label can be max 63 characters
		labelLen := min(63, remaining)
		if labelLen <= 0 {
			break
		}

		// Generate label
		for i := range labelLen {
			result.WriteByte(byte('a' + ((labelCount*labelLen + i) % 26)))
		}

		remaining -= labelLen
		labelCount++
	}

	return result.String()
}
