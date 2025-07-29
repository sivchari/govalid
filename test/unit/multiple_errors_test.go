package unit_test

import (
	"errors"
	"testing"

	"github.com/go-playground/validator/v10"

	"github.com/sivchari/govalid/test"
	govaliderrors "github.com/sivchari/govalid/validation/errors"
)

func TeserrtMultipleErrorsValidation(t *testing.T) {
	tests := []struct {
		name           string
		data           test.MultipleErrors
		expectedErrors int
		expectError    error
	}{
		{
			name:           "multiple errors",
			data:           test.MultipleErrors{TooLong: "error"},
			expectedErrors: 2,
		},
		{
			name:           "single error",
			data:           test.MultipleErrors{URL: "single error", TooLong: "1"},
			expectedErrors: 1,
		},
		{
			name:           "no errors",
			data:           test.MultipleErrors{URL: "single error", TooLong: "1"},
			expectedErrors: 0,
		},
	}

	for _, tt := range tests {
		t.Run("govalid_"+tt.name, func(t *testing.T) {
			err := test.ValidateMultipleErrors(&tt.data)
			if tt.expectedErrors <= 0 {
				if err != nil {
					t.Errorf("execution expected no errors returned %#v", err)
				}
			}

			var validationErrors govaliderrors.ValidationErrors
			if errors.As(err, &validationErrors) {
				if len(validationErrors) != tt.expectedErrors {
					t.Errorf("execution produced more errors than expected %#v", validationErrors)
				}
			}
		})

		t.Run("go-playground_"+tt.name, func(t *testing.T) {
			validate := validator.New()
			err := validate.Struct(tt.data)
			if tt.expectedErrors <= 0 {
				if err != nil {
					t.Errorf("execution expected no errors returned %#v", err)
				}
			}

			if err == nil && tt.expectedErrors <= 0 {
				t.Errorf("expected error but got nil")
			}
		})
	}
}
