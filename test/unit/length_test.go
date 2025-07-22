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
			name:        "exactly at limit",
			data:        test.Length{Name: "1234567"},
			expectError: false,
		},
		{
			name:        "limit minus one",
			data:        test.Length{Name: "123456"},
			expectError: true,
		},
		{
			name:        "limit plus one",
			data:        test.Length{Name: "12345678"},
			expectError: true,
		},
		{
			name:        "empty string",
			data:        test.Length{Name: ""},
			expectError: true,
		},
		{
			name:        "unicode characters - valid",
			data:        test.Length{Name: "🔥🔥🔥🔥🔥🔥🔥"},
			expectError: false,
		},
		{
			name:        "unicode characters - invalid",
			data:        test.Length{Name: "🔥🔥🔥🔥🔥🔥"},
			expectError: true,
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
