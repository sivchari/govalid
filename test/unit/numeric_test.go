package unit

import (
	"testing"

	"github.com/go-playground/validator/v10"

	"github.com/sivchari/govalid/test"
)

func TestNumericValidation(t *testing.T) {
	tests := []struct {
		name        string
		data        test.Numeric
		expectError bool
	}{
		{"valid_zero", test.Numeric{Number: "0"}, false},
		{"valid_digits", test.Numeric{Number: "1234567890"}, false},
		{"valid_leading_zeros", test.Numeric{Number: "000000"}, false},
		{"valid_nines", test.Numeric{Number: "999999"}, false},

		// Invalid cases
		{"empty", test.Numeric{Number: ""}, true},
		{"alpha_numeric", test.Numeric{Number: "123abc"}, true},
		{"decimal", test.Numeric{Number: "12.34"}, true},
		{"leading_space", test.Numeric{Number: " 123"}, true},
		{"trailing_space", test.Numeric{Number: "123 "}, true},
		{"dash", test.Numeric{Number: "12-34"}, true},
		{"underscore", test.Numeric{Number: "12_34"}, true},
		{"plus", test.Numeric{Number: "12+34"}, true},
		{"letters_only", test.Numeric{Number: "abc"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test govalid
			err := test.ValidateNumeric(&tt.data)
			hasError := err != nil
			if hasError != tt.expectError {
				t.Errorf("govalid: expected error=%v, got error=%v (err: %v)", tt.expectError, hasError, err)
			}

			// Test go-playground/validator for comparison
			validate := validator.New()
			validate.RegisterValidation("numeric", func(fl validator.FieldLevel) bool {
				val := fl.Field().String()
				for _, ch := range val {
					if ch < '0' || ch > '9' {
						return false
					}
				}
				return val != ""
			})

			validatorStruct := struct {
				Number string `validate:"numeric"`
			}{Number: tt.data.Number}

			err = validate.Struct(validatorStruct)
			hasError = err != nil
			if hasError != tt.expectError {
				t.Errorf("go-playground/validator: expected error=%v, got error=%v (err: %v)", tt.expectError, hasError, err)
			}
		})
	}
}
