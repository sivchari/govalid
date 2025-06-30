package unit

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sivchari/govalid/test"
)

// Test data types for GT validation behavior comparison
type GTTestData test.TestGT

func TestGTValidation(t *testing.T) {
	validate := validator.New()

	tests := []struct {
		name        string
		data        GTTestData
		expectError bool
		description string
	}{
		{
			name:        "valid gt value",
			data:        GTTestData{Age: 150},
			expectError: false,
			description: "age 150 is greater than limit 100",
		},
		{
			name:        "equal to limit",
			data:        GTTestData{Age: 100},
			expectError: true,
			description: "age equal to limit should be invalid",
		},
		{
			name:        "less than limit",
			data:        GTTestData{Age: 50},
			expectError: true,
			description: "age less than limit should be invalid",
		},
		// Boundary value tests
		{
			name:        "one more than limit",
			data:        GTTestData{Age: 101},
			expectError: false,
			description: "age 101 is exactly one more than limit 100",
		},
		{
			name:        "one less than limit",
			data:        GTTestData{Age: 99},
			expectError: true,
			description: "age 99 is exactly one less than limit 100",
		},
		{
			name:        "zero value",
			data:        GTTestData{Age: 0},
			expectError: true,
			description: "zero is less than limit 100",
		},
		{
			name:        "negative value",
			data:        GTTestData{Age: -5},
			expectError: true,
			description: "negative values are less than limit 100",
		},
		{
			name:        "minimum valid value",
			data:        GTTestData{Age: 101},
			expectError: false,
			description: "minimum value that satisfies gt constraint",
		},
		{
			name:        "maximum invalid value",
			data:        GTTestData{Age: 100},
			expectError: true,
			description: "maximum value that violates gt constraint",
		},
		{
			name:        "large valid value",
			data:        GTTestData{Age: 1000000},
			expectError: false,
			description: "very large value should be valid",
		},
		{
			name:        "minimum int",
			data:        GTTestData{Age: -2147483648}, // int32 min for compatibility
			expectError: true,
			description: "minimum integer value is less than limit",
		},
		{
			name:        "just above boundary",
			data:        GTTestData{Age: 100 + 1},
			expectError: false,
			description: "computed boundary plus one",
		},
		{
			name:        "just at boundary",
			data:        GTTestData{Age: 100 + 0},
			expectError: true,
			description: "computed boundary exact",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test govalid
			govalidErr := test.ValidateGT((*test.GT)(&tt.data))
			govalidHasError := govalidErr != nil

			// Test go-playground/validator
			playgroundErr := validate.Struct(&tt.data)
			playgroundHasError := playgroundErr != nil

			// Compare results
			if govalidHasError != tt.expectError {
				t.Errorf("govalid: expected error=%v, got error=%v (%v) - %s", tt.expectError, govalidHasError, govalidErr, tt.description)
			}
			if playgroundHasError != tt.expectError {
				t.Errorf("go-playground: expected error=%v, got error=%v (%v) - %s", tt.expectError, playgroundHasError, playgroundErr, tt.description)
			}
			if govalidHasError != playgroundHasError {
				t.Errorf("behavior mismatch: govalid=%v, playground=%v - %s", govalidHasError, playgroundHasError, tt.description)
			}
		})
	}
}