package unit

import (
	"testing"

	"github.com/sivchari/govalid/test"
)

func TestAlphaValidation(t *testing.T) {
	tests := []struct {
		name        string
		data        test.Alpha
		expectError bool
	}{
		{"valid_case-1", test.Alpha{FirstName: "John"}, false},
		{"valid_case-2", test.Alpha{FirstName: "John"}, false},
		{"valid_case-3", test.Alpha{FirstName: "JOHn"}, false},
		{"valid_case-4", test.Alpha{FirstName: "JOHN"}, false},

		// Invalid cases
		{"invalid_case-1", test.Alpha{FirstName: "John1"}, true},
		{"invalid_case-2", test.Alpha{FirstName: "John Mayor"}, true},
		{"invalid_case-3", test.Alpha{FirstName: "123"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := test.ValidateAlpha(&tt.data)
			hasError := err != nil

			if hasError != tt.expectError {
				t.Errorf("expected error: %v, got error: %v (%v)", tt.expectError, hasError, err)
			}
		})
	}
}
