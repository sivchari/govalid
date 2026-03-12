package unit

import (
	"context"
	"testing"

	"github.com/sivchari/govalid"
	"github.com/sivchari/govalid/test"
)

func TestValidator(t *testing.T) {
	tests := []struct {
		name        string
		validator   govalid.Validator
		expectError bool
	}{
		// Alpha tests
		{"Alpha_valid", &test.Alpha{FirstName: "John", LastName: "Doe", CountryCode: "US"}, false},
		{"Alpha_invalid", &test.Alpha{FirstName: "John1", LastName: "Doe", CountryCode: "US"}, true},

		// GT tests
		{"GT_valid", &test.GT{Age: 101}, false},
		{"GT_invalid", &test.GT{Age: 50}, true},

		// GTE tests
		{"GTE_valid", &test.GTE{Age: 18}, false},
		{"GTE_invalid", &test.GTE{Age: 17}, true},

		// MaxLength tests
		{"MaxLength_valid", &test.MaxLength{Name: "short"}, false},
		{"MaxLength_invalid", &test.MaxLength{Name: "this is a very long name that definitely exceeds fifty characters which is the limit"}, true},

		// Required tests
		{"Required_valid", &test.Required{Name: "John", Age: 25, Items: []string{}}, false},
		{"Required_invalid", &test.Required{Name: ""}, true},

		// Email tests
		{"Email_valid", &test.Email{Email: "user@example.com"}, false},
		{"Email_invalid", &test.Email{Email: "invalid-email"}, true},

		// LT tests
		{"LT_valid", &test.LT{Age: 9}, false},
		{"LT_invalid", &test.LT{Age: 18}, true},

		// LTE tests
		{"LTE_valid", &test.LTE{Age: 55}, false},
		{"LTE_invalid", &test.LTE{Age: 101}, true},

		// Numeric tests
		{"Numeric_valid", &test.Numeric{Number: "123"}, false},
		{"Numeric_invalid", &test.Numeric{Number: "not-a-number"}, true},

		// IPV4 tests
		{"IPV4_valid", &test.IPV4{IP: "192.168.1.1"}, false},
		{"IPV4_invalid", &test.IPV4{IP: "invalid"}, true},

		// IPV6 tests
		{"IPV6_valid", &test.IPV6{IP: "2001:0db8:85a3:0000:0000:8a2e:0370:7334"}, false},
		{"IPV6_invalid", &test.IPV6{IP: "invalid"}, true},

		// URL tests
		{"URL_valid", &test.URL{URL: "https://example.com"}, false},
		{"URL_invalid", &test.URL{URL: "not-a-url"}, true},

		// UUID tests
		{"UUID_valid", &test.UUID{UUID: "550e8400-e29b-41d4-a716-446655440000"}, false},
		{"UUID_invalid", &test.UUID{UUID: "invalid"}, true},

		// Length tests
		{"Length_valid", &test.Length{Name: "1234567"}, false},
		{"Length_invalid", &test.Length{Name: "short"}, true},

		// MinLength tests
		{"MinLength_valid", &test.MinLength{Name: "abc"}, false},
		{"MinLength_invalid", &test.MinLength{Name: "ab"}, true},

		// MinItems tests
		{"MinItems_valid", &test.MinItems{Items: []string{"a", "b"}, Metadata: map[string]string{"k": "v"}, ChanField: func() chan int { ch := make(chan int, 1); ch <- 1; return ch }()}, false},
		{"MinItems_invalid", &test.MinItems{Items: []string{"a"}}, true},

		// MaxItems tests
		{"MaxItems_valid", &test.MaxItems{Items: []string{"a"}}, false},
		{"MaxItems_invalid", &test.MaxItems{Items: []string{"a", "b", "c", "d", "e", "f"}}, true},

		// Enum tests
		{"Enum_valid", &test.Enum{Role: "admin", Level: 1, UserRole: "manager", Priority: 10}, false},
		{"Enum_invalid", &test.Enum{Role: "invalid", Level: 1, UserRole: "manager", Priority: 10}, true},

		// CEL tests
		{"CEL_valid", &test.CEL{Age: 18, Name: "John", Score: 1.0, IsActive: true}, false},
		{"CEL_invalid", &test.CEL{Age: 17, Name: "John", Score: 1.0, IsActive: true}, true},

		// CELCrossField tests
		{"CELCrossField_valid", &test.CELCrossField{Price: 10, MaxPrice: 20, Quantity: 2, Budget: 100}, false},
		{"CELCrossField_invalid", &test.CELCrossField{Price: 30, MaxPrice: 20, Quantity: 2, Budget: 100}, true},

		// MultipleErrors tests
		{"MultipleErrors_valid", &test.MultipleErrors{URL: "something", TooLong: "a"}, false},
		{"MultipleErrors_invalid", &test.MultipleErrors{URL: "", TooLong: "toolong"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.validator.Validate()
			hasError := err != nil

			if hasError != tt.expectError {
				t.Errorf("expected error: %v, got error: %v (%v)", tt.expectError, hasError, err)
			}
		})
	}
}

func TestValidatorContext(t *testing.T) {
	tests := []struct {
		name        string
		validator   govalid.ContextValidator
		expectError bool
	}{
		// Alpha tests
		{"Alpha_valid", &test.Alpha{FirstName: "John", LastName: "Doe", CountryCode: "US"}, false},
		{"Alpha_invalid", &test.Alpha{FirstName: "John1", LastName: "Doe", CountryCode: "US"}, true},

		// GT tests
		{"GT_valid", &test.GT{Age: 101}, false},
		{"GT_invalid", &test.GT{Age: 50}, true},

		// GTE tests
		{"GTE_valid", &test.GTE{Age: 18}, false},
		{"GTE_invalid", &test.GTE{Age: 17}, true},

		// MaxLength tests
		{"MaxLength_valid", &test.MaxLength{Name: "short"}, false},
		{"MaxLength_invalid", &test.MaxLength{Name: "this is a very long name that definitely exceeds fifty characters which is the limit"}, true},

		// Required tests
		{"Required_valid", &test.Required{Name: "John", Age: 25, Items: []string{}}, false},
		{"Required_invalid", &test.Required{Name: ""}, true},

		// Email tests
		{"Email_valid", &test.Email{Email: "user@example.com"}, false},
		{"Email_invalid", &test.Email{Email: "invalid-email"}, true},

		// LT tests
		{"LT_valid", &test.LT{Age: 9}, false},
		{"LT_invalid", &test.LT{Age: 18}, true},

		// LTE tests
		{"LTE_valid", &test.LTE{Age: 55}, false},
		{"LTE_invalid", &test.LTE{Age: 101}, true},

		// Numeric tests
		{"Numeric_valid", &test.Numeric{Number: "123"}, false},
		{"Numeric_invalid", &test.Numeric{Number: "not-a-number"}, true},

		// IPV4 tests
		{"IPV4_valid", &test.IPV4{IP: "192.168.1.1"}, false},
		{"IPV4_invalid", &test.IPV4{IP: "invalid"}, true},

		// IPV6 tests
		{"IPV6_valid", &test.IPV6{IP: "2001:0db8:85a3:0000:0000:8a2e:0370:7334"}, false},
		{"IPV6_invalid", &test.IPV6{IP: "invalid"}, true},

		// URL tests
		{"URL_valid", &test.URL{URL: "https://example.com"}, false},
		{"URL_invalid", &test.URL{URL: "not-a-url"}, true},

		// UUID tests
		{"UUID_valid", &test.UUID{UUID: "550e8400-e29b-41d4-a716-446655440000"}, false},
		{"UUID_invalid", &test.UUID{UUID: "invalid"}, true},

		// Length tests
		{"Length_valid", &test.Length{Name: "1234567"}, false},
		{"Length_invalid", &test.Length{Name: "short"}, true},

		// MinLength tests
		{"MinLength_valid", &test.MinLength{Name: "abc"}, false},
		{"MinLength_invalid", &test.MinLength{Name: "ab"}, true},

		// MinItems tests
		{"MinItems_valid", &test.MinItems{Items: []string{"a", "b"}, Metadata: map[string]string{"k": "v"}, ChanField: func() chan int { ch := make(chan int, 1); ch <- 1; return ch }()}, false},
		{"MinItems_invalid", &test.MinItems{Items: []string{"a"}}, true},

		// MaxItems tests
		{"MaxItems_valid", &test.MaxItems{Items: []string{"a"}}, false},
		{"MaxItems_invalid", &test.MaxItems{Items: []string{"a", "b", "c", "d", "e", "f"}}, true},

		// Enum tests
		{"Enum_valid", &test.Enum{Role: "admin", Level: 1, UserRole: "manager", Priority: 10}, false},
		{"Enum_invalid", &test.Enum{Role: "invalid", Level: 1, UserRole: "manager", Priority: 10}, true},

		// CEL tests
		{"CEL_valid", &test.CEL{Age: 18, Name: "John", Score: 1.0, IsActive: true}, false},
		{"CEL_invalid", &test.CEL{Age: 17, Name: "John", Score: 1.0, IsActive: true}, true},

		// CELCrossField tests
		{"CELCrossField_valid", &test.CELCrossField{Price: 10, MaxPrice: 20, Quantity: 2, Budget: 100}, false},
		{"CELCrossField_invalid", &test.CELCrossField{Price: 30, MaxPrice: 20, Quantity: 2, Budget: 100}, true},

		// MultipleErrors tests
		{"MultipleErrors_valid", &test.MultipleErrors{URL: "something", TooLong: "a"}, false},
		{"MultipleErrors_invalid", &test.MultipleErrors{URL: "", TooLong: "toolong"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.validator.ValidateContext(context.Background())
			hasError := err != nil

			if hasError != tt.expectError {
				t.Errorf("expected error: %v, got error: %v (%v)", tt.expectError, hasError, err)
			}
		})
	}
}
