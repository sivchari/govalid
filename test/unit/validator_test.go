package unit

import (
	"testing"

	"github.com/sivchari/govalid"
	"github.com/sivchari/govalid/test"
)

func TestValidator(t *testing.T) {
	tests := []struct {
		name        string
		validator   govalid.Validator
		expectError bool
	}{
		// Alpha tests
		{"Alpha_valid", &test.Alpha{FirstName: "John", LastName: "Doe", CountryCode: "US"}, false},
		{"Alpha_invalid", &test.Alpha{FirstName: "John1", LastName: "Doe", CountryCode: "US"}, true},

		// GT tests
		{"GT_valid", &test.GT{Age: 101}, false},
		{"GT_invalid", &test.GT{Age: 50}, true},

		// GTE tests
		{"GTE_valid", &test.GTE{Age: 18}, false},
		{"GTE_invalid", &test.GTE{Age: 17}, true},

		// MaxLength tests
		{"MaxLength_valid", &test.MaxLength{Name: "short"}, false},
		{"MaxLength_invalid", &test.MaxLength{Name: "this is a very long name that definitely exceeds fifty characters which is the limit"}, true},

		// Required tests
		{"Required_valid", &test.Required{Name: "John", Age: 25, Items: []string{}}, false},
		{"Required_invalid", &test.Required{Name: ""}, true},

		// Email tests
		{"Email_valid", &test.Email{Email: "user@example.com"}, false},
		{"Email_invalid", &test.Email{Email: "invalid-email"}, true},

		// LT tests
		{"LT_valid", &test.LT{Age: 9}, false},
		{"LT_invalid", &test.LT{Age: 18}, true},

		// LTE tests
		{"LTE_valid", &test.LTE{Age: 55}, false},
		{"LTE_invalid", &test.LTE{Age: 101}, true},

		// Numeric tests
		{"Numeric_valid", &test.Numeric{Number: "123"}, false},
		{"Numeric_invalid", &test.Numeric{Number: "not-a-number"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.validator.Validate()
			hasError := err != nil

			if hasError != tt.expectError {
				t.Errorf("expected error: %v, got error: %v (%v)", tt.expectError, hasError, err)
			}
		})
	}
}
