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
		{
			name:        "valid alphanum",
			data:        test.Alphanum{ProductCode: "product2024", SerialNumber: "ABC123", Username: "username"},
			expectError: false,
		},
		{
			name:        "invalid alphanum",
			data:        test.Alphanum{ProductCode: "product_2024", SerialNumber: "ABC-123", Username: "username@@@"},
			expectError: true,
		},
		{
			name:        "empty string",
			data:        test.Alphanum{ProductCode: "", SerialNumber: "", Username: ""},
			expectError: true,
		},
		{
			name:        "unicode characters - invalid",
			data:        test.Alphanum{ProductCode: "你", SerialNumber: "٣", Username: "🔥🔥🔥🔥🔥🔥🔥"},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			govalidErr := test.ValidateAlphanum(&tt.data)
			govalidHasError := govalidErr != nil
			if govalidHasError != tt.expectError {
				t.Errorf("govalid: expected error=%v, got error=%v (err: %v)", tt.expectError, govalidHasError, govalidErr)
			}
		})
	}
}
