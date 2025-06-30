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
		// Core boundary value tests around limit=50
		{
			name:        "empty string",
			data:        MaxLengthTestData{Name: ""},
			expectError: false,
			description: "empty string should be valid",
		},
		{
			name:        "limit minus one",
			data:        MaxLengthTestData{Name: strings.Repeat("x", 49)}, // 49 chars
			expectError: false,
			description: "one character under limit",
		},
		{
			name:        "exactly at limit",
			data:        MaxLengthTestData{Name: strings.Repeat("x", 50)}, // 50 chars
			expectError: false,
			description: "exactly at 50 character limit",
		},
		{
			name:        "limit plus one",
			data:        MaxLengthTestData{Name: strings.Repeat("x", 51)}, // 51 chars
			expectError: true,
			description: "one character over limit",
		},
		// Unicode boundary test
		{
			name:        "unicode at limit",
			data:        MaxLengthTestData{Name: strings.Repeat("🔥", 50)}, // 50 unicode chars
			expectError: false,
			description: "unicode characters at limit",
		},
		{
			name:        "unicode over limit",
			data:        MaxLengthTestData{Name: strings.Repeat("🔥", 51)}, // 51 unicode chars
			expectError: true,
			description: "unicode characters over limit",
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

			// Compare results - both should now use rune count consistently
			if govalidHasError != tt.expectError {
				t.Errorf("govalid: expected error=%v, got error=%v (%v) - %s (runes: %d, bytes: %d)", tt.expectError, govalidHasError, govalidErr, tt.description, len([]rune(tt.data.Name)), len(tt.data.Name))
			}
			if playgroundHasError != tt.expectError {
				t.Errorf("go-playground: expected error=%v, got error=%v (%v) - %s (runes: %d, bytes: %d)", tt.expectError, playgroundHasError, playgroundErr, tt.description, len([]rune(tt.data.Name)), len(tt.data.Name))
			}
			if govalidHasError != playgroundHasError {
				t.Errorf("behavior mismatch: govalid=%v, playground=%v - %s (runes: %d, bytes: %d)", govalidHasError, playgroundHasError, tt.description, len([]rune(tt.data.Name)), len(tt.data.Name))
			}
		})
	}
}