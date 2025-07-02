package unit

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sivchari/govalid/test"
)

func TestLTEValidation(t *testing.T) {
	validate := validator.New()

	tests := []struct {
		name        string
		data        test.LTE
		expectError bool
	}{
		{
			name:        "below maximum",
			data:        test.LTE{Age: 50},
			expectError: false,
		},
		{
			name:        "exactly at maximum",
			data:        test.LTE{Age: 100},
			expectError: false,
		},
		{
			name:        "above maximum",
			data:        test.LTE{Age: 150},
			expectError: true,
		},
		{
			name:        "zero value",
			data:        test.LTE{Age: 0},
			expectError: false,
		},
		{
			name:        "negative value",
			data:        test.LTE{Age: -5},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test govalid
			govalidErr := test.ValidateLTE(&tt.data)
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