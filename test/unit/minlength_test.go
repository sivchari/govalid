package unit

import (
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sivchari/govalid/test"
)

func TestMinLengthValidation(t *testing.T) {
	validate := validator.New()

	tests := []struct {
		name        string
		data        test.MinLength
		expectError bool
	}{
		{
			name:        "empty string",
			data:        test.MinLength{Name: ""},
			expectError: true,
		},
		{
			name:        "limit minus one",
			data:        test.MinLength{Name: strings.Repeat("x", 2)},
			expectError: true,
		},
		{
			name:        "exactly at limit",
			data:        test.MinLength{Name: strings.Repeat("x", 3)},
			expectError: false,
		},
		{
			name:        "limit plus one",
			data:        test.MinLength{Name: strings.Repeat("x", 4)},
			expectError: false,
		},
		{
			name:        "unicode at limit",
			data:        test.MinLength{Name: "🔥🔥🔥"},
			expectError: false,
		},
		{
			name:        "unicode below limit",
			data:        test.MinLength{Name: "🔥🔥"},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test govalid
			govalidErr := test.ValidateMinLength(&tt.data)
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
