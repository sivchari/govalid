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
		// String enum tests
		{
			name:        "valid string enum - admin",
			data:        test.Enum{Role: "admin", Level: 1, UserRole: "manager", Priority: 10},
			expectError: false,
		},
		{
			name:        "valid string enum - user",
			data:        test.Enum{Role: "user", Level: 2, UserRole: "developer", Priority: 20},
			expectError: false,
		},
		{
			name:        "valid string enum - guest",
			data:        test.Enum{Role: "guest", Level: 3, UserRole: "tester", Priority: 30},
			expectError: false,
		},
		{
			name:        "invalid string enum",
			data:        test.Enum{Role: "invalid", Level: 1, UserRole: "manager", Priority: 10},
			expectError: true,
		},
		{
			name:        "empty string enum",
			data:        test.Enum{Role: "", Level: 1, UserRole: "manager", Priority: 10},
			expectError: true,
		},
		// Numeric enum tests
		{
			name:        "invalid numeric enum - level 0",
			data:        test.Enum{Role: "admin", Level: 0, UserRole: "manager", Priority: 10},
			expectError: true,
		},
		{
			name:        "invalid numeric enum - level 4",
			data:        test.Enum{Role: "admin", Level: 4, UserRole: "manager", Priority: 10},
			expectError: true,
		},
		// Custom string type enum tests
		{
			name:        "invalid custom string enum",
			data:        test.Enum{Role: "admin", Level: 1, UserRole: "invalid", Priority: 10},
			expectError: true,
		},
		// Custom numeric type enum tests
		{
			name:        "invalid custom numeric enum",
			data:        test.Enum{Role: "admin", Level: 1, UserRole: "manager", Priority: 5},
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