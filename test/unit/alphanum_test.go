package unit

import (
	"testing"

	"github.com/sivchari/govalid/test"
)

func TestAlphanumValidation(t *testing.T) {
	tests := []struct {
		name        string
		data        test.Alphanum
		expectError bool
	}{
		// Valid cases
		{"letters", test.Alphanum{ProductCode: "ABCDEF", SerialNumber: "abcdef", Username: "Product"}, false},
		{"letters and digits", test.Alphanum{ProductCode: "ABC123", SerialNumber: "product2024", Username: "X1"}, false},
		{"digits only", test.Alphanum{ProductCode: "12345", SerialNumber: "67890", Username: "0"}, false},

		// Invalid cases
		{"with hyphen", test.Alphanum{ProductCode: "ABC-123", SerialNumber: "abc", Username: "abc"}, true},
		{"with underscore", test.Alphanum{ProductCode: "abc", SerialNumber: "product_2024", Username: "abc"}, true},
		{"with dot", test.Alphanum{ProductCode: "abc", SerialNumber: "abc", Username: "12.34"}, true},
		{"with leading space", test.Alphanum{ProductCode: " ABC123", SerialNumber: "abc", Username: "abc"}, true},
		{"with inner space", test.Alphanum{ProductCode: "abc", SerialNumber: "ABC 123", Username: "abc"}, true},
		{"empty string", test.Alphanum{ProductCode: "", SerialNumber: "abc", Username: "abc"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := test.ValidateAlphanum(&tt.data)
			hasError := err != nil

			if hasError != tt.expectError {
				t.Errorf("expected error: %v, got error: %v (%v)", tt.expectError, hasError, err)
			}
		})
	}
}
