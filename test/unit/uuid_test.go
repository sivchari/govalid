package unit

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sivchari/govalid/test"
)

func TestUUIDValidation(t *testing.T) {
	tests := []struct {
		name        string
		data        test.UUID
		expectError bool
	}{
		// Valid UUIDs (v1, v2, v3, v4, v5) - lowercase only to match go-playground
		{"valid_v1", test.UUID{UUID: "550e8400-e29b-11d4-a716-446655440000"}, false},
		{"valid_v2", test.UUID{UUID: "000003e8-2363-21ef-b200-325096b39f47"}, false},
		{"valid_v3", test.UUID{UUID: "6ba7b810-9dad-31d1-80b4-00c04fd430c8"}, false},
		{"valid_v4", test.UUID{UUID: "6ba7b811-9dad-41d1-80b4-00c04fd430c8"}, false},
		{"valid_v5", test.UUID{UUID: "6ba7b812-9dad-51d1-80b4-00c04fd430c8"}, false},
		{"valid_lowercase", test.UUID{UUID: "123e4567-e89b-12d3-a456-426614174000"}, false},
		
		// Invalid UUIDs
		{"empty", test.UUID{UUID: ""}, true},
		{"too_short", test.UUID{UUID: "123e4567-e89b-12d3-a456"}, true},
		{"too_long", test.UUID{UUID: "123e4567-e89b-12d3-a456-4266141740001"}, true},
		{"no_hyphens", test.UUID{UUID: "123e4567e89b12d3a456426614174000"}, true},
		{"wrong_hyphens", test.UUID{UUID: "123e4567e89b-12d3-a456-426614174000"}, true},
		{"invalid_chars", test.UUID{UUID: "123g4567-e89b-12d3-a456-426614174000"}, true},  // 'g' is not hex
		{"spaces", test.UUID{UUID: "123e4567 e89b-12d3-a456-426614174000"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test govalid
			err := test.ValidateUUID(&tt.data)
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