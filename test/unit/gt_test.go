package unit

import (
	"testing"

	"github.com/go-playground/validator/v10"

	"github.com/sivchari/govalid/test"
)

func TestGTValidation(t *testing.T) {
	validate := validator.New()

	tests := []struct {
		name        string
		data        test.GT
		expectError bool
	}{
		{
			name:        "valid",
			data:        test.GT{Age: 150},
			expectError: false,
		},
		{
			name:        "limit minus one",
			data:        test.GT{Age: 99},
			expectError: true,
		},
		{
			name:        "equal to limit",
			data:        test.GT{Age: 100},
			expectError: true,
		},
		{
			name:        "limit plus one",
			data:        test.GT{Age: 101},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test govalid
			govalidErr := test.ValidateGT(&tt.data)
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
