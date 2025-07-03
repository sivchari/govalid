package unit

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sivchari/govalid/test"
)

func TestURLValidation(t *testing.T) {
	tests := []struct {
		name        string
		data        test.URL
		expectError bool
	}{
		// Valid URLs
		{"valid_http", test.URL{URL: "http://example.com"}, false},
		{"valid_https", test.URL{URL: "https://example.com"}, false},
		{"valid_with_port", test.URL{URL: "https://example.com:8080"}, false},
		{"valid_with_path", test.URL{URL: "https://example.com/path/to/resource"}, false},
		{"valid_with_query", test.URL{URL: "https://example.com?param=value"}, false},
		{"valid_with_fragment", test.URL{URL: "https://example.com#section"}, false},
		{"valid_complex", test.URL{URL: "https://subdomain.example.com:8080/path?query=value#fragment"}, false},
		{"valid_subdomain", test.URL{URL: "https://api.example.com"}, false},
		{"valid_with_dash", test.URL{URL: "https://my-site.example.com"}, false},
		
		// Invalid URLs
		{"empty", test.URL{URL: ""}, true},
		{"no_protocol", test.URL{URL: "example.com"}, true},
		{"no_domain", test.URL{URL: "https://"}, true},
		{"spaces", test.URL{URL: "https://example .com"}, true},
		{"just_protocol", test.URL{URL: "https"}, true},
		{"malformed", test.URL{URL: "https//example.com"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test govalid
			err := test.ValidateURL(&tt.data)
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