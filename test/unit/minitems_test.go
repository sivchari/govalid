package unit

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sivchari/govalid/test"
)

func TestMinItemsValidation(t *testing.T) {
	validate := validator.New()

	tests := []struct {
		name        string
		data        test.MinItems
		expectError bool
	}{
		{
			name:        "empty slice",
			data:        test.MinItems{Items: []string{}},
			expectError: true,
		},
		{
			name:        "limit minus one",
			data:        test.MinItems{Items: []string{"a"}},
			expectError: true,
		},
		{
			name:        "exactly at limit",
			data:        test.MinItems{Items: []string{"a", "b"}},
			expectError: false,
		},
		{
			name:        "limit plus one",
			data:        test.MinItems{Items: []string{"a", "b", "c"}},
			expectError: false,
		},
		{
			name:        "nil slice",
			data:        test.MinItems{Items: nil},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test govalid
			govalidErr := test.ValidateMinItems(&tt.data)
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