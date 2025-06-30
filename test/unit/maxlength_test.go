package unit

import (
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sivchari/govalid/test"
)

// Test data types for MaxLength validation behavior comparison
type MaxLengthTestData test.TestMaxLength

func TestMaxLengthValidation(t *testing.T) {
	validate := validator.New()

	tests := []struct {
		name        string
		data        MaxLengthTestData
		expectError bool
		description string
	}{
		{
			name:        "valid length",
			data:        MaxLengthTestData{Name: "Valid Name"},
			expectError: false,
			description: "normal string within limit",
		},
		{
			name:        "empty string",
			data:        MaxLengthTestData{Name: ""},
			expectError: false,
			description: "empty string should be valid for maxlength",
		},
		{
			name:        "exactly at limit",
			data:        MaxLengthTestData{Name: strings.Repeat("x", 50)}, // 50 chars
			expectError: false,
			description: "string exactly at the 50 character limit",
		},
		{
			name:        "exceeds limit by one",
			data:        MaxLengthTestData{Name: strings.Repeat("x", 51)}, // 51 chars
			expectError: true,
			description: "string exceeding limit by exactly one character",
		},
		{
			name:        "far exceeds limit",
			data:        MaxLengthTestData{Name: strings.Repeat("x", 100)}, // 100 chars
			expectError: true,
			description: "string far exceeding the limit",
		},
		// Boundary value tests
		{
			name:        "one character",
			data:        MaxLengthTestData{Name: "x"},
			expectError: false,
			description: "single character string",
		},
		{
			name:        "limit minus one",
			data:        MaxLengthTestData{Name: strings.Repeat("x", 49)}, // 49 chars
			expectError: false,
			description: "string one character under the limit",
		},
		{
			name:        "limit plus two",
			data:        MaxLengthTestData{Name: strings.Repeat("x", 52)}, // 52 chars
			expectError: true,
			description: "string two characters over the limit",
		},
		{
			name:        "unicode characters - byte count at limit",
			data:        MaxLengthTestData{Name: strings.Repeat("a", 50)}, // 50 ascii chars = 50 bytes
			expectError: false,
			description: "ascii characters using byte-based counting at limit",
		},
		{
			name:        "unicode characters - byte count over limit",
			data:        MaxLengthTestData{Name: strings.Repeat("🔥", 13)}, // 13 unicode chars = 52 bytes (over limit)
			expectError: true,
			description: "unicode characters exceeding byte limit (govalid uses byte count, playground uses rune count)",
		},
		{
			name:        "unicode - known behavior difference",
			data:        MaxLengthTestData{Name: strings.Repeat("🔥", 12)}, // 12 unicode chars = 48 bytes
			expectError: false,
			description: "unicode characters under byte limit",
		},
		{
			name:        "whitespace at limit",
			data:        MaxLengthTestData{Name: strings.Repeat(" ", 50)}, // 50 spaces
			expectError: false,
			description: "whitespace characters at the limit",
		},
		{
			name:        "whitespace over limit",
			data:        MaxLengthTestData{Name: strings.Repeat(" ", 51)}, // 51 spaces
			expectError: true,
			description: "whitespace characters over the limit",
		},
		{
			name:        "newlines and tabs at limit",
			data:        MaxLengthTestData{Name: strings.Repeat("\n\t", 25)}, // 50 chars
			expectError: false,
			description: "newlines and tabs at the limit",
		},
		{
			name:        "very long string",
			data:        MaxLengthTestData{Name: strings.Repeat("verylongstring", 100)}, // 1400 chars
			expectError: true,
			description: "extremely long string should be invalid",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test govalid
			govalidErr := test.ValidateMaxLength((*test.MaxLength)(&tt.data))
			govalidHasError := govalidErr != nil

			// Test go-playground/validator
			playgroundErr := validate.Struct(&tt.data)
			playgroundHasError := playgroundErr != nil

			// Compare results - handle known Unicode behavior differences
			if strings.Contains(tt.name, "unicode characters - byte count over limit") {
				// Known difference: govalid uses byte count, playground uses rune count for Unicode
				if !govalidHasError {
					t.Errorf("govalid: expected to treat Unicode byte overflow as invalid, got valid")
				}
				if playgroundHasError {
					t.Errorf("go-playground: expected to treat Unicode rune count as valid, got invalid")
				}
			} else {
				// For all other cases, expect identical behavior
				if govalidHasError != tt.expectError {
					t.Errorf("govalid: expected error=%v, got error=%v (%v) - %s (length: %d)", tt.expectError, govalidHasError, govalidErr, tt.description, len(tt.data.Name))
				}
				if playgroundHasError != tt.expectError {
					t.Errorf("go-playground: expected error=%v, got error=%v (%v) - %s (length: %d)", tt.expectError, playgroundHasError, playgroundErr, tt.description, len(tt.data.Name))
				}
				if govalidHasError != playgroundHasError {
					t.Errorf("behavior mismatch: govalid=%v, playground=%v - %s (length: %d)", govalidHasError, playgroundHasError, tt.description, len(tt.data.Name))
				}
			}
		})
	}
}