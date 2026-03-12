package errors

import (
	"errors"
	"testing"
)

func TestValidationError_Error(t *testing.T) {
	t.Parallel()

	err := ValidationError{
		Path:   "Name",
		Type:   "required",
		Value:  "",
		Reason: "field Name is required",
	}

	got := err.Error()
	want := `field Name with value  has failed validation required because field Name is required`

	if got != want {
		t.Errorf("Error() = %q, want %q", got, want)
	}
}

func TestValidationError_Is(t *testing.T) {
	t.Parallel()

	err := ValidationError{
		Path:   "Name",
		Type:   "required",
		Reason: "field Name is required",
	}

	tests := []struct {
		name   string
		target error
		want   bool
	}{
		{
			"same error",
			ValidationError{Path: "Name", Type: "required", Reason: "field Name is required"},
			true,
		},
		{
			"different path",
			ValidationError{Path: "Email", Type: "required", Reason: "field Name is required"},
			false,
		},
		{
			"different type",
			ValidationError{Path: "Name", Type: "email", Reason: "field Name is required"},
			false,
		},
		{
			"different reason",
			ValidationError{Path: "Name", Type: "required", Reason: "different reason"},
			false,
		},
		{
			"non-validation error",
			errors.New("some error"),
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := err.Is(tt.target); got != tt.want {
				t.Errorf("Is(%v) = %v, want %v", tt.target, got, tt.want)
			}
		})
	}
}

func TestValidationErrors_Error(t *testing.T) {
	t.Parallel()

	errs := ValidationErrors{
		{Path: "Name", Type: "required", Reason: "field Name is required"},
		{Path: "Email", Type: "email", Reason: "field Email is not valid email"},
	}

	got := errs.Error()

	if got == "" {
		t.Error("Error() returned empty string")
	}
}

func TestValidationErrors_Is(t *testing.T) {
	t.Parallel()

	errs := ValidationErrors{
		{Path: "Name", Type: "required", Reason: "field Name is required"},
		{Path: "Email", Type: "email", Reason: "field Email is not valid email"},
	}

	tests := []struct {
		name   string
		target error
		want   bool
	}{
		{
			"matching first error",
			ValidationError{Path: "Name", Type: "required", Reason: "field Name is required"},
			true,
		},
		{
			"matching second error",
			ValidationError{Path: "Email", Type: "email", Reason: "field Email is not valid email"},
			true,
		},
		{
			"no match",
			ValidationError{Path: "Age", Type: "gt", Reason: "field Age must be greater"},
			false,
		},
		{
			"non-validation error",
			errors.New("some error"),
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := errs.Is(tt.target); got != tt.want {
				t.Errorf("Is(%v) = %v, want %v", tt.target, got, tt.want)
			}
		})
	}
}
