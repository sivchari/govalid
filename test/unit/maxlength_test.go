package unit

import (
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sivchari/govalid/test"
)

func TestMaxLengthValidation(t *testing.T) {
	validate := validator.New()

	tests := []struct {
		name        string
		data        test.MaxLength
		expectError bool
	}{
		{
			name:        "empty string",
			data:        test.MaxLength{Name: ""},
			expectError: false,
		},
		{
			name:        "limit minus one",
			data:        test.MaxLength{Name: strings.Repeat("x", 49)},
			expectError: false,
		},
		{
			name:        "exactly at limit",
			data:        test.MaxLength{Name: strings.Repeat("x", 50)},
			expectError: false,
		},
		{
			name:        "limit plus one",
			data:        test.MaxLength{Name: strings.Repeat("x", 51)},
			expectError: true,
		},
		{
			name:        "unicode over limit",
			data:        test.MaxLength{Name: strings.Repeat("🔥", 51)},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test govalid
			govalidErr := test.ValidateMaxLength(&tt.data)
			govalidHasError := govalidErr != nil

			// Test go-playground/validator
			playgroundErr := validate.Struct(&tt.data)
			playgroundHasError := playgroundErr != nil

			// Compare results
			if govalidHasError != tt.expectError {
				t.Errorf("govalid: expected error=%v, got error=%v (%v)", tt.expectError, govalidHasError, govalidErr)
			}
			if playgroundHasError != tt.expectError {
				t.Errorf("go-playground: expected error=%v, got error=%v (%v)", tt.expectError, playgroundHasError, playgroundErr)
			}
			if govalidHasError != playgroundHasError {
				t.Errorf("behavior mismatch: govalid=%v, playground=%v", govalidHasError, playgroundHasError)
			}
		})
	}
}
