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

// TestURLValidationDifferences tests cases where govalid and go-playground/validator behave differently.
// This documents the design philosophy differences between the two libraries.
func TestMultipleErrorsValidationDifferences(t *testing.T) {
	tests := []struct {
		name string
		// MultipleErrors stucts
		url     string
		tooLong string
		// ---
		govalidValid      bool
		goplaygroundValid bool
		reason            string
	}{
		{
			name:              "You're taking too long",
			url:               "",
			tooLong:           "too long",
			govalidValid:      false,
			goplaygroundValid: false,
			reason:            "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test govalid
			data := test.MultipleErrors{URL: tt.url, TooLong: tt.tooLong}
			govalidErr := test.ValidateMultipleErrors(&data)
			govalidValid := govalidErr == nil

			if govalidValid != tt.govalidValid {
				t.Errorf("govalid: expected valid=%v, got valid=%v for %s", tt.govalidValid, govalidValid, tt.url)
			}

			t.Errorf("%#v", govalidErr)
			t.Errorf("%s", govalidErr.Error())

			// Test go-playground
			validate := validator.New()
			goplaygroundErr := validate.Struct(&data)
			goplaygroundValid := goplaygroundErr == nil

			t.Errorf("%#v", goplaygroundErr)
			t.Errorf("%s", goplaygroundErr.Error())

			if goplaygroundValid != tt.goplaygroundValid {
				t.Errorf("go-playground: expected valid=%v, got valid=%v for %s",
					tt.goplaygroundValid, goplaygroundValid, tt.url)
			}

			// Log the difference for documentation
			if govalidValid != goplaygroundValid {
				t.Logf("Design difference for %s: %s", tt.url, tt.reason)
			}
		})
	}
}
