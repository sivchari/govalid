package unit

import (
	"testing"

	"github.com/go-playground/validator/v10"

	"github.com/sivchari/govalid/test"
)

func TestGTEValidation(t *testing.T) {
	validate := validator.New()

	tests := []struct {
		name        string
		data        test.GTE
		expectError bool
	}{
		{
			name:        "below minimum",
			data:        test.GTE{Age: 17},
			expectError: true,
		},
		{
			name:        "exactly at minimum",
			data:        test.GTE{Age: 18},
			expectError: false,
		},
		{
			name:        "above minimum",
			data:        test.GTE{Age: 25},
			expectError: false,
		},
		{
			name:        "zero value",
			data:        test.GTE{Age: 0},
			expectError: true,
		},
		{
			name:        "negative value",
			data:        test.GTE{Age: -5},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test govalid
			govalidErr := test.ValidateGTE(&tt.data)
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
