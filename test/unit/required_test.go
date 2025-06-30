package unit

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sivchari/govalid/test"
)

// Test data types for required validation behavior comparison
type RequiredTestData test.TestRequired

func TestRequiredValidation(t *testing.T) {
	validate := validator.New()

	tests := []struct {
		name        string
		data        RequiredTestData
		expectError bool
		description string
	}{
		{
			name:        "valid required fields",
			data:        RequiredTestData{Name: "John", Age: 25, Items: []string{"item1"}},
			expectError: false,
			description: "all required fields have valid values",
		},
		{
			name:        "empty name",
			data:        RequiredTestData{Name: "", Age: 25, Items: []string{"item1"}},
			expectError: true,
			description: "name is empty string",
		},
		{
			name:        "zero age",
			data:        RequiredTestData{Name: "John", Age: 0, Items: []string{"item1"}},
			expectError: true,
			description: "age is zero value",
		},
		{
			name:        "nil items",
			data:        RequiredTestData{Name: "John", Age: 25, Items: nil},
			expectError: true,
			description: "items slice is nil",
		},
		{
			name:        "empty items",
			data:        RequiredTestData{Name: "John", Age: 25, Items: []string{}},
			expectError: true,
			description: "govalid treats empty slice as invalid, playground treats as valid - behavior difference",
		},
		// Boundary value tests
		{
			name:        "single character name",
			data:        RequiredTestData{Name: "A", Age: 1, Items: []string{"x"}},
			expectError: false,
			description: "minimum valid string length",
		},
		{
			name:        "minimum positive age",
			data:        RequiredTestData{Name: "John", Age: 1, Items: []string{"item"}},
			expectError: false,
			description: "minimum positive integer value",
		},
		{
			name:        "negative age",
			data:        RequiredTestData{Name: "John", Age: -1, Items: []string{"item"}},
			expectError: false,
			description: "negative values are considered non-zero",
		},
		{
			name:        "single item in slice",
			data:        RequiredTestData{Name: "John", Age: 25, Items: []string{"single"}},
			expectError: false,
			description: "slice with exactly one element",
		},
		{
			name:        "whitespace only name",
			data:        RequiredTestData{Name: " ", Age: 25, Items: []string{"item"}},
			expectError: false,
			description: "whitespace is considered a valid non-empty string",
		},
		{
			name:        "all fields at boundary - valid",
			data:        RequiredTestData{Name: "x", Age: 1, Items: []string{"y"}},
			expectError: false,
			description: "all fields at minimum valid boundary",
		},
		{
			name:        "all fields at boundary - invalid",
			data:        RequiredTestData{Name: "", Age: 0, Items: []string{}},
			expectError: true,
			description: "all fields at invalid boundary",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test govalid
			govalidErr := test.ValidateRequired((*test.Required)(&tt.data))
			govalidHasError := govalidErr != nil

			// Test go-playground/validator
			playgroundErr := validate.Struct(&tt.data)
			playgroundHasError := playgroundErr != nil

			// Compare results - allow for known behavior difference in empty slice handling
			if tt.name == "empty items" {
				// Known difference: govalid treats empty slices as invalid, playground treats as valid
				if !govalidHasError {
					t.Errorf("govalid: expected to treat empty slice as invalid, got valid")
				}
				if playgroundHasError {
					t.Errorf("go-playground: expected to treat empty slice as valid, got invalid")
				}
			} else {
				// For all other cases, expect identical behavior
				if govalidHasError != tt.expectError {
					t.Errorf("govalid: expected error=%v, got error=%v (%v) - %s", tt.expectError, govalidHasError, govalidErr, tt.description)
				}
				if playgroundHasError != tt.expectError {
					t.Errorf("go-playground: expected error=%v, got error=%v (%v) - %s", tt.expectError, playgroundHasError, playgroundErr, tt.description)
				}
				if govalidHasError != playgroundHasError {
					t.Errorf("behavior mismatch: govalid=%v, playground=%v - %s", govalidHasError, playgroundHasError, tt.description)
				}
			}
		})
	}
}