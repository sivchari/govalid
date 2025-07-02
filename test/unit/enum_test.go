package unit

import (
	"testing"

	"github.com/sivchari/govalid/test"
)

func TestEnumValidation(t *testing.T) {
	// Note: go-playground/validator doesn't have a direct enum validator,
	// so we test only govalid behavior

	tests := []struct {
		name        string
		data        test.Enum
		expectError bool
	}{
		{
			name:        "valid enum value - admin",
			data:        test.Enum{Role: "admin"},
			expectError: false,
		},
		{
			name:        "valid enum value - user",
			data:        test.Enum{Role: "user"},
			expectError: false,
		},
		{
			name:        "valid enum value - guest",
			data:        test.Enum{Role: "guest"},
			expectError: false,
		},
		{
			name:        "invalid enum value",
			data:        test.Enum{Role: "superuser"},
			expectError: true,
		},
		{
			name:        "empty string",
			data:        test.Enum{Role: ""},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test govalid
			govalidErr := test.ValidateEnum(&tt.data)
			govalidHasError := govalidErr != nil

			if govalidHasError != tt.expectError {
				t.Errorf("govalid: expected error=%v, got error=%v (%v)", tt.expectError, govalidHasError, govalidErr)
			}
		})
	}
}