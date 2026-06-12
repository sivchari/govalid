package validationhelper

import "testing"

func TestIsValidAlphanum(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"lowercase", "abc", true},
		{"uppercase", "ABC", true},
		{"digits", "123", true},
		{"mixed case", "aBcDeF", true},
		{"letters and digits", "abc123", true},
		{"single letter", "a", true},
		{"single digit", "1", true},
		{"empty string", "", false},
		{"with space", "ab c", false},
		{"with special char", "abc!", false},
		{"with underscore", "abc_def", false},
		{"with hyphen", "abc-def", false},
		{"with dot", "12.34", false},
		{"leading space", " abc", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := IsValidAlphanum(tt.input)
			if got != tt.want {
				t.Errorf("IsValidAlphanum(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
