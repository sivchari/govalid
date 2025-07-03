package unit

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sivchari/govalid/test"
)

func TestEmailValidation(t *testing.T) {
	tests := []struct {
		name        string
		data        test.Email
		expectError bool
	}{
		{"valid_simple", test.Email{Email: "test@example.com"}, false},
		{"valid_with_dots", test.Email{Email: "first.last@example.com"}, false},
		{"valid_with_numbers", test.Email{Email: "user123@domain.org"}, false},
		{"valid_with_plus", test.Email{Email: "user+tag@example.com"}, false},
		{"valid_subdomain", test.Email{Email: "user@sub.example.com"}, false},
		{"valid_with_special_chars", test.Email{Email: "user.name+tag@example-domain.com"}, false},
		
		// Invalid cases  
		{"empty", test.Email{Email: ""}, true},
		{"no_at_symbol", test.Email{Email: "userexample.com"}, true},
		{"multiple_at_symbols", test.Email{Email: "user@@example.com"}, true},
		{"no_domain", test.Email{Email: "user@"}, true},
		{"no_local_part", test.Email{Email: "@example.com"}, true},
		{"spaces", test.Email{Email: "user @example.com"}, true},
		{"invalid_characters", test.Email{Email: "user@exam ple.com"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test govalid
			err := test.ValidateEmail(&tt.data)
			hasError := err != nil
			if hasError != tt.expectError {
				t.Errorf("govalid: expected error=%v, got error=%v (err: %v)", tt.expectError, hasError, err)
			}

			// Test go-playground/validator for comparison
			validate := validator.New()
			err = validate.Struct(&tt.data)
			hasError = err != nil
			if hasError != tt.expectError {
				t.Errorf("go-playground/validator: expected error=%v, got error=%v (err: %v)", tt.expectError, hasError, err)
			}
		})
	}
}