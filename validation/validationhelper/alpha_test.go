package validationhelper

import "testing"

func TestIsValidAlpha(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{"lowercase", "abc", true},
		{"uppercase", "ABC", true},
		{"mixed case", "aBcDeF", true},
		{"single char", "a", true},
		{"empty string", "", true},
		{"with digit", "abc1", false},
		{"with space", "ab c", false},
		{"only digits", "123", false},
		{"with special char", "abc!", false},
		{"with underscore", "abc_def", false},
		{"with hyphen", "abc-def", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := IsValidAlpha(tt.input)
			if got != tt.want {
				t.Errorf("IsValidAlpha(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
