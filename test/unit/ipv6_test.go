package unit

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sivchari/govalid/test"
)

func TestIPV6Validation(t *testing.T) {
	validate := validator.New()

	tests := []struct {
		name        string
		data        test.IPV6
		expectError bool
	}{
		{
			name:        "ipv6 address",
			data:        test.IPV6{IP: "2001:0db8:85a3:0000:0000:8a2e:0370:7334"},
			expectError: false,
		},
		{
			name:        "full ipv6 address",
			data:        test.IPV6{IP: "2001:0db8:85a3:0000:0000:8a2e:0370:7334"},
			expectError: false,
		},
		{
			name:        "ipv6 localhost",
			data:        test.IPV6{IP: "::1"},
			expectError: false,
		},
		{
			name:        "ipv6 unspecified address",
			data:        test.IPV6{IP: "::"},
			expectError: false,
		},
		{
			name:        "ipv6 with trailing compression",
			data:        test.IPV6{IP: "2001:db8::"},
			expectError: false,
		},
		{
			name:        "ipv4 address",
			data:        test.IPV6{IP: "192.168.1.0"},
			expectError: true,
		},
		{
			name:        "empty string",
			data:        test.IPV6{IP: ""},
			expectError: true,
		},
		{
			name:        "contains invalid characters",
			data:        test.IPV6{IP: "2001:0db8:85a3:xyz0:0000:8a2e:0370:7334"},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test govalid
			govalidErr := test.ValidateIPV6(&tt.data)
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
