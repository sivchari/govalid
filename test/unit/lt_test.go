package unit

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sivchari/govalid/test"
)

// Test data types for LT validation behavior comparison
type LTTestData test.TestLT

func TestLTValidation(t *testing.T) {
	validate := validator.New()

	tests := []struct {
		name        string
		data        LTTestData
		expectError bool
		description string
	}{
		{
			name:        "valid lt value",
			data:        LTTestData{Age: 5},
			expectError: false,
			description: "age 5 is less than limit 10",
		},
		{
			name:        "equal to limit",
			data:        LTTestData{Age: 10},
			expectError: true,
			description: "age equal to limit should be invalid",
		},
		{
			name:        "greater than limit",
			data:        LTTestData{Age: 15},
			expectError: true,
			description: "age greater than limit should be invalid",
		},
		// Boundary value tests
		{
			name:        "one less than limit",
			data:        LTTestData{Age: 9},
			expectError: false,
			description: "age 9 is exactly one less than limit 10",
		},
		{
			name:        "one more than limit",
			data:        LTTestData{Age: 11},
			expectError: true,
			description: "age 11 is exactly one more than limit 10",
		},
		{
			name:        "zero value",
			data:        LTTestData{Age: 0},
			expectError: false,
			description: "zero is less than limit 10",
		},
		{
			name:        "negative value",
			data:        LTTestData{Age: -5},
			expectError: false,
			description: "negative values are less than limit 10",
		},
		{
			name:        "minimum int",
			data:        LTTestData{Age: -2147483648}, // int32 min for compatibility
			expectError: false,
			description: "minimum integer value is less than limit",
		},
		{
			name:        "maximum valid value",
			data:        LTTestData{Age: 9},
			expectError: false,
			description: "maximum value that satisfies lt constraint",
		},
		{
			name:        "minimum invalid value",
			data:        LTTestData{Age: 10},
			expectError: true,
			description: "minimum value that violates lt constraint",
		},
		{
			name:        "large invalid value",
			data:        LTTestData{Age: 1000000},
			expectError: true,
			description: "very large value should still be invalid",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test govalid
			govalidErr := test.ValidateLT((*test.LT)(&tt.data))
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