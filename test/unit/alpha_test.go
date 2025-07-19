package unit

import (
	"github.com/sivchari/govalid/test"
	"testing"
)

func TestAlphaValidation(t *testing.T) {
	tests := []struct {
		name        string
		data        test.Alpha
		expectError bool
	}{
		{"valid_case", test.Alpha{FirstName: "valid_value"}, false},
		{"boundary_minus_one", test.Alpha{FirstName: "boundary-1"}, false},
		{"exactly_at_boundary", test.Alpha{FirstName: "boundary"}, false},
		{"boundary_plus_one", test.Alpha{FirstName: "boundary+1"}, true},
		{"invalid_case", test.Alpha{FirstName: "invalid_value"}, true},
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
