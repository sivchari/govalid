package unit_test

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
		{
			name: "valid_https_url",
			data: test.URL{
				WebsiteURL: "https://example.com",
			},
			expectError: false,
		},
		{
			name: "valid_http_url",
			data: test.URL{
				WebsiteURL: "http://example.com",
			},
			expectError: false,
		},
		{
			name: "valid_url_with_path",
			data: test.URL{
				WebsiteURL: "https://example.com/path/to/resource",
			},
			expectError: false,
		},
		{
			name: "valid_url_with_query",
			data: test.URL{
				WebsiteURL: "https://example.com?key=value&foo=bar",
			},
			expectError: false,
		},
		{
			name: "valid_url_with_port",
			data: test.URL{
				WebsiteURL: "https://example.com:8080",
			},
			expectError: false,
		},
		{
			name: "valid_ftp_url",
			data: test.URL{
				WebsiteURL: "ftp://files.example.com",
			},
			expectError: false,
		},
		{
			name: "invalid_no_scheme",
			data: test.URL{
				WebsiteURL: "example.com",
			},
			expectError: true,
		},
		{
			name: "invalid_no_host",
			data: test.URL{
				WebsiteURL: "https://",
			},
			expectError: true,
		},
		{
			name: "invalid_empty_string",
			data: test.URL{
				WebsiteURL: "",
			},
			expectError: true,
		},
		{
			name: "invalid_spaces",
			data: test.URL{
				WebsiteURL: "https://example .com",
			},
			expectError: true,
		},
		{
			name: "valid_mailto",
			data: test.URL{
				WebsiteURL: "mailto:test@example.com",
			},
			expectError: false,
		},
		{
			name: "valid_git_url",
			data: test.URL{
				WebsiteURL: "git://github.com/user/repo.git",
			},
			expectError: false,
		},
		{
			name: "valid_ssh_url",
			data: test.URL{
				WebsiteURL: "ssh://user@server.com:22",
			},
			expectError: false,
		},

		// Edge cases for schemes without host
		{
			name: "invalid_mailto_empty",
			data: test.URL{
				WebsiteURL: "mailto:",
			},
			expectError: true,
		},
		{
			name: "valid_mailto_simple",
			data: test.URL{
				WebsiteURL: "mailto:test@example.com",
			},
			expectError: false,
		},
		{
			name: "invalid_data_empty",
			data: test.URL{
				WebsiteURL: "data:",
			},
			expectError: true,
		},
		{
			name: "valid_data_simple",
			data: test.URL{
				WebsiteURL: "data:text/plain,Hello",
			},
			expectError: false,
		},
		{
			name: "valid_file_path",
			data: test.URL{
				WebsiteURL: "file:/path/to/file",
			},
			expectError: false,
		},

		// Edge cases for schemes with host
		{
			name: "invalid_http_no_slashes",
			data: test.URL{
				WebsiteURL: "http:",
			},
			expectError: true,
		},
		{
			name: "invalid_http_one_slash",
			data: test.URL{
				WebsiteURL: "http:/",
			},
			expectError: true,
		},
		{
			name: "invalid_http_two_slashes_no_host",
			data: test.URL{
				WebsiteURL: "http://",
			},
			expectError: true,
		},
		{
			name: "valid_ipv6_brackets",
			data: test.URL{
				WebsiteURL: "http://[::1]",
			},
			expectError: false,
		},
		{
			name: "valid_numeric_host",
			data: test.URL{
				WebsiteURL: "http://192.168.1.1",
			},
			expectError: false,
		},

		// Scheme validation edge cases
		{
			name: "invalid_scheme_with_invalid_char",
			data: test.URL{
				WebsiteURL: "ht@tp://example.com",
			},
			expectError: true,
		},
		{
			name: "invalid_no_colon",
			data: test.URL{
				WebsiteURL: "httpexample.com",
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run("govalid_"+tt.name, func(t *testing.T) {
			err := test.ValidateURL(&tt.data)
			if tt.expectError && err == nil {
				t.Errorf("expected error but got nil")
			}
			if !tt.expectError && err != nil {
				t.Errorf("expected no error but got %v", err)
			}
		})

		t.Run("go-playground_"+tt.name, func(t *testing.T) {
			validate := validator.New()
			err := validate.Struct(tt.data)
			if tt.expectError && err == nil {
				t.Errorf("expected error but got nil")
			}
			if !tt.expectError && err != nil {
				t.Errorf("expected no error but got %v", err)
			}
		})
	}
}

// TestURLValidationDifferences tests cases where govalid and go-playground/validator behave differently.
// This documents the design philosophy differences between the two libraries.
func TestURLValidationDifferences(t *testing.T) {
	tests := []struct {
		name              string
		url               string
		govalidValid      bool
		goplaygroundValid bool
		reason            string
	}{
		{
			name:              "host_starts_with_dot",
			url:               "http://.example.com",
			govalidValid:      false,
			goplaygroundValid: true,
			reason:            "govalid validates host format, go-playground is more lenient",
		},
		{
			name:              "unknown_scheme",
			url:               "svn+ssh://example.com",
			govalidValid:      false,
			goplaygroundValid: true,
			reason:            "govalid validates against known schemes, go-playground accepts any scheme",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test govalid
			data := test.URL{WebsiteURL: tt.url}
			govalidErr := test.ValidateURL(&data)
			govalidValid := govalidErr == nil

			if govalidValid != tt.govalidValid {
				t.Errorf("govalid: expected valid=%v, got valid=%v for %s",
					tt.govalidValid, govalidValid, tt.url)
			}

			// Test go-playground
			validate := validator.New()
			goplaygroundErr := validate.Struct(&data)
			goplaygroundValid := goplaygroundErr == nil

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
