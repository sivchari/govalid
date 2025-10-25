package unit

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sivchari/govalid/test"
)

func TestIPV4Validation(t *testing.T) {
	validate := validator.New()

	tests := []struct {
		name        string
		data        test.IPV4
		expectError bool
	}{
		{
			name:        "ipv4 address",
			data:        test.IPV4{IP: "192.168.1.0"},
			expectError: false,
		},
		{
			name:        "ipv4 localhost",
			data:        test.IPV4{IP: "127.0.0.1"},
			expectError: false,
		},
		{
			name:        "ipv4 with zeros",
			data:        test.IPV4{IP: "0.0.0.0"},
			expectError: false,
		},
		{
			name:        "ipv4 max values",
			data:        test.IPV4{IP: "255.255.255.255"},
			expectError: false,
		},
		{
			name:        "ipv6 address",
			data:        test.IPV4{IP: "2001:0db8:85a3:0000:0000:8a2e:0370:7334"},
			expectError: true,
		},
		{
			name:        "octet exceeds 255",
			data:        test.IPV4{IP: "192.168.1.256"},
			expectError: true,
		},
		{
			name:        "missing octet",
			data:        test.IPV4{IP: "192.168.1"},
			expectError: true,
		},
		{
			name:        "too many octets",
			data:        test.IPV4{IP: "192.168.1.1.1"},
			expectError: true,
		},
		{
			name:        "empty string",
			data:        test.IPV4{IP: ""},
			expectError: true,
		},
		{
			name:        "ipv4 with port",
			data:        test.IPV4{IP: "192.168.1.1:8080"},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test govalid
			govalidErr := test.ValidateIPV4(&tt.data)
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
