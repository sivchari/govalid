package unit_test

import (
	"errors"
	"testing"

	"github.com/sivchari/govalid/test"
)

func TestValidateMultipleErrors_Is(t *testing.T) {
	testCases := []struct {
		name       string
		input      *test.MultipleErrors
		expectErr  bool
		targets    []error // errors to check for with errors.Is
		nonTargets []error // errors that should NOT be found
	}{
		{
			name: "valid case",
			input: &test.MultipleErrors{
				URL:     "http://example.com",
				TooLong: "a",
			},
			expectErr: false,
		},
		{
			name: "error: URL is required",
			input: &test.MultipleErrors{
				URL:     "",
				TooLong: "a",
			},
			expectErr:  true,
			targets:    []error{test.ErrMultipleErrorsURLRequiredValidation},
			nonTargets: []error{test.ErrMultipleErrorsTooLongMaxLengthValidation},
		},
		{
			name: "error: TooLong is too long",
			input: &test.MultipleErrors{
				URL:     "http://example.com",
				TooLong: "ab",
			},
			expectErr:  true,
			targets:    []error{test.ErrMultipleErrorsTooLongMaxLengthValidation},
			nonTargets: []error{test.ErrMultipleErrorsURLRequiredValidation},
		},
		{
			name: "error: both URL required and TooLong is too long",
			input: &test.MultipleErrors{
				URL:     "",
				TooLong: "ab",
			},
			expectErr: true,
			targets: []error{
				test.ErrMultipleErrorsURLRequiredValidation,
				test.ErrMultipleErrorsTooLongMaxLengthValidation,
			},
		},
		{
			name:      "error: nil input",
			input:     nil,
			expectErr: true,
			targets:   []error{test.ErrNilMultipleErrors},
			nonTargets: []error{
				test.ErrMultipleErrorsURLRequiredValidation,
				test.ErrMultipleErrorsTooLongMaxLengthValidation,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := test.ValidateMultipleErrors(tc.input)

			if (err != nil) != tc.expectErr {
				t.Fatalf("expected error: %v, got: %v", tc.expectErr, err)
			}

			if !tc.expectErr {
				return
			}

			for _, target := range tc.targets {
				if !errors.Is(err, target) {
					t.Errorf("expected error to be '%v', but it was not", target)
				}
			}

			for _, nonTarget := range tc.nonTargets {
				if errors.Is(err, nonTarget) {
					t.Errorf("expected error not to be '%v', but it was", nonTarget)
				}
			}
		})
	}
}
