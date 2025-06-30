package unit

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sivchari/govalid/test"
)

// Test data types for validation behavior comparison
type RequiredTestData test.TestRequired
type LTTestData test.TestLT
type GTTestData test.TestGT
type MaxLengthTestData test.TestMaxLength

func TestRequiredValidation(t *testing.T) {
	validate := validator.New()

	tests := []struct {
		name        string
		data        RequiredTestData
		expectError bool
	}{
		{
			name:        "valid required fields",
			data:        RequiredTestData{Name: "John", Age: 25, Items: []string{"item1"}},
			expectError: false,
		},
		{
			name:        "empty name",
			data:        RequiredTestData{Name: "", Age: 25, Items: []string{"item1"}},
			expectError: true,
		},
		{
			name:        "zero age",
			data:        RequiredTestData{Name: "John", Age: 0, Items: []string{"item1"}},
			expectError: true,
		},
		{
			name:        "nil items",
			data:        RequiredTestData{Name: "John", Age: 25, Items: nil},
			expectError: true,
		},
		{
			name:        "empty items",
			data:        RequiredTestData{Name: "John", Age: 25, Items: []string{}},
			expectError: true, // govalid treats empty slice as invalid, playground treats as valid - behavior difference
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

func TestLTValidation(t *testing.T) {
	validate := validator.New()

	tests := []struct {
		name        string
		data        LTTestData
		expectError bool
	}{
		{
			name:        "valid lt value",
			data:        LTTestData{Age: 5},
			expectError: false,
		},
		{
			name:        "equal to limit",
			data:        LTTestData{Age: 10},
			expectError: true,
		},
		{
			name:        "greater than limit",
			data:        LTTestData{Age: 15},
			expectError: true,
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

func TestGTValidation(t *testing.T) {
	validate := validator.New()

	tests := []struct {
		name        string
		data        GTTestData
		expectError bool
	}{
		{
			name:        "valid gt value",
			data:        GTTestData{Age: 150},
			expectError: false,
		},
		{
			name:        "equal to limit",
			data:        GTTestData{Age: 100},
			expectError: true,
		},
		{
			name:        "less than limit",
			data:        GTTestData{Age: 50},
			expectError: true,
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

func TestMaxLengthValidation(t *testing.T) {
	validate := validator.New()

	tests := []struct {
		name        string
		data        MaxLengthTestData
		expectError bool
	}{
		{
			name:        "valid length",
			data:        MaxLengthTestData{Name: "Valid Name"},
			expectError: false,
		},
		{
			name:        "empty string",
			data:        MaxLengthTestData{Name: ""},
			expectError: false,
		},
		{
			name:        "exactly at limit",
			data:        MaxLengthTestData{Name: "12345678901234567890123456789012345678901234567890"}, // 50 chars
			expectError: false,
		},
		{
			name:        "exceeds limit",
			data:        MaxLengthTestData{Name: "123456789012345678901234567890123456789012345678901"}, // 51 chars
			expectError: true,
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