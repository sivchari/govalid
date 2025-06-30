package unit

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sivchari/govalid/test"
)

func TestRequiredValidation(t *testing.T) {
	validate := validator.New()

	tests := []struct {
		name        string
		data        test.Required
		expectError bool
	}{
		{
			name:        "valid",
			data:        test.Required{Name: "John", Age: 25, Items: []string{"item"}},
			expectError: false,
		},
		{
			name:        "empty name",
			data:        test.Required{Name: "", Age: 25, Items: []string{"item"}},
			expectError: true,
		},
		{
			name:        "zero age",
			data:        test.Required{Name: "John", Age: 0, Items: []string{"item"}},
			expectError: true,
		},
		{
			name:        "empty slice",
			data:        test.Required{Name: "John", Age: 25, Items: []string{}},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test govalid
			govalidErr := test.ValidateRequired(&tt.data)
			govalidHasError := govalidErr != nil

			// Test go-playground/validator
			playgroundErr := validate.Struct(&tt.data)
			playgroundHasError := playgroundErr != nil

			// Compare results - allow for known behavior difference in empty slice handling
			if tt.name == "empty slice" {
				// Known difference: govalid treats empty slices as invalid (correct), playground treats as valid
				if !govalidHasError {
					t.Errorf("govalid: expected to treat empty slice as invalid, got valid")
				}
				if playgroundHasError {
					t.Errorf("go-playground: expected to treat empty slice as valid, got invalid")
				}
			} else {
				// For all other cases, expect identical behavior
				if govalidHasError != tt.expectError {
					t.Errorf("govalid: expected error=%v, got error=%v (%v)", tt.expectError, govalidHasError, govalidErr)
				}
				if playgroundHasError != tt.expectError {
					t.Errorf("go-playground: expected error=%v, got error=%v (%v)", tt.expectError, playgroundHasError, playgroundErr)
				}
				if govalidHasError != playgroundHasError {
					t.Errorf("behavior mismatch: govalid=%v, playground=%v", govalidHasError, playgroundHasError)
				}
			}
		})
	}
}