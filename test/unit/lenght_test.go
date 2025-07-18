package unit

import (
	"testing"

	"github.com/sivchari/govalid/test"
)

func TestLengthValidation(t *testing.T) {
	tests := []struct {
		name        string
		data        test.Length
		expectError bool
	}{
		{
			name:        "valid postal code and phone number",
			data:        test.Length{PostalCode: "1234567", PhoneNumber: "1234567890"},
			expectError: false,
		},
		{
			name:        "postal code too long",
			data:        test.Length{PostalCode: "12345678", PhoneNumber: "1234567890"},
			expectError: true,
		},
		{
			name:        "postal code too short",
			data:        test.Length{PostalCode: "123456", PhoneNumber: "1234567890"},
			expectError: true,
		},
		{
			name:        "phone number too long",
			data:        test.Length{PostalCode: "1234567", PhoneNumber: "12345678901"},
			expectError: true,
		},
		{
			name:        "phone number too short",
			data:        test.Length{PostalCode: "1234567", PhoneNumber: "123456789"},
			expectError: true,
		},
		{
			name:        "both invalid",
			data:        test.Length{PostalCode: "123", PhoneNumber: "123"},
			expectError: true,
		},
		{
			name:        "unicode characters - valid",
			data:        test.Length{PostalCode: "1234567", PhoneNumber: "🔥🔥🔥🔥🔥🔥🔥🔥🔥🔥"},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test govalid
			govalidErr := test.ValidateLength(&tt.data)
			govalidHasError := govalidErr != nil

			// Compare results (only testing govalid since go-playground/validator
			// doesn't have a direct equivalent to exact length validation)
			if govalidHasError != tt.expectError {
				t.Errorf("govalid: expected error=%v, got error=%v (%v)", tt.expectError, govalidHasError, govalidErr)
			}
		})
	}
}
