package unit

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sivchari/govalid/test"
)

func TestMaxItemsValidation(t *testing.T) {
	validate := validator.New()

	tests := []struct {
		name        string
		data        test.MaxItems
		expectError bool
	}{
		{
			name:        "empty slice",
			data:        test.MaxItems{Items: []string{}},
			expectError: false,
		},
		{
			name:        "limit minus one",
			data:        test.MaxItems{Items: []string{"a", "b", "c", "d"}},
			expectError: false,
		},
		{
			name:        "exactly at limit",
			data:        test.MaxItems{Items: []string{"a", "b", "c", "d", "e"}},
			expectError: false,
		},
		{
			name:        "limit plus one",
			data:        test.MaxItems{Items: []string{"a", "b", "c", "d", "e", "f"}},
			expectError: true,
		},
		{
			name:        "nil slice",
			data:        test.MaxItems{Items: nil},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test govalid
			govalidErr := test.ValidateMaxItems(&tt.data)
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