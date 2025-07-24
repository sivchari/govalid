package validationhelper

import (
	"testing"
)

func TestIsValidUUID(t *testing.T) {
	t.Run("valid_uuids", testValidUUIDs)
	t.Run("invalid_uuids", testInvalidUUIDs)
}

func testValidUUIDs(t *testing.T) {
	tests := []struct {
		name     string
		uuid     string
		expected bool
	}{
		// Valid UUIDs - Version 1 (time-based)
		{"version_1_lowercase", "550e8400-e29b-11d4-a716-446655440000", true},
		{"version_1_uppercase", "550E8400-E29B-11D4-A716-446655440000", true},
		{"version_1_mixed", "550e8400-E29B-11d4-A716-446655440000", true},

		// Valid UUIDs - Version 2 (DCE Security)
		{"version_2", "550e8400-e29b-21d4-a716-446655440000", true},

		// Valid UUIDs - Version 3 (MD5)
		{"version_3", "550e8400-e29b-31d4-a716-446655440000", true},
		{"version_3_variant_8", "550e8400-e29b-31d4-8716-446655440000", true},
		{"version_3_variant_9", "550e8400-e29b-31d4-9716-446655440000", true},
		{"version_3_variant_b", "550e8400-e29b-31d4-b716-446655440000", true},

		// Valid UUIDs - Version 4 (random)
		{"version_4", "550e8400-e29b-41d4-a716-446655440000", true},
		{"version_4_real_example", "f47ac10b-58cc-4372-a567-0e02b2c3d479", true},
		{"version_4_variant_lowercase_a", "550e8400-e29b-41d4-a716-446655440000", true},
		{"version_4_variant_uppercase_a", "550e8400-e29b-41d4-A716-446655440000", true},
		{"version_4_variant_lowercase_b", "550e8400-e29b-41d4-b716-446655440000", true},
		{"version_4_variant_uppercase_b", "550e8400-e29b-41d4-B716-446655440000", true},

		// Valid UUIDs - Version 5 (SHA-1)
		{"version_5", "550e8400-e29b-51d4-a716-446655440000", true},
		{"version_5_real_example", "2ed6657d-e927-568b-95e1-2665a8aea6a2", true},

		// Valid UUIDs - Special cases
		{"nil_uuid", "00000000-0000-0000-0000-000000000000", true},
		{"all_f", "ffffffff-ffff-ffff-ffff-ffffffffffff", true},
		{"numeric_only", "12345678-1234-1234-8234-123456789012", true},
		{"letters_only", "abcdefab-abcd-1bcd-abcd-abcdefabcdef", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidUUID(tt.uuid)
			if result != tt.expected {
				t.Errorf("IsValidUUID(%q) = %v, expected %v", tt.uuid, result, tt.expected)
			}
		})
	}
}

func testInvalidUUIDs(t *testing.T) {
	tests := []struct {
		name     string
		uuid     string
		expected bool
	}{
		// Invalid UUIDs - Wrong length
		{"empty_string", "", false},
		{"too_short", "550e8400-e29b-41d4-a716-44665544000", false},
		{"too_long", "550e8400-e29b-41d4-a716-4466554400000", false},
		{"missing_group", "550e8400-e29b-41d4-446655440000", false},
		{"extra_group", "550e8400-e29b-41d4-a716-4466-55440000", false},

		// Invalid UUIDs - Wrong hyphen positions
		{"no_hyphens", "550e8400e29b41d4a716446655440000", false},
		{"wrong_hyphen_pos_1", "550e840-0e29b-41d4-a716-446655440000", false},
		{"wrong_hyphen_pos_2", "550e8400-e29b41d4-a716-446655440000", false},
		{"wrong_hyphen_pos_3", "550e8400-e29b-41d4a716-446655440000", false},
		{"wrong_hyphen_pos_4", "550e8400-e29b-41d4-a716446655440000", false},
		{"extra_hyphen", "550e8400-e29b-41d4-a716-4466-55440000", false},
		{"hyphen_at_start", "-550e8400-e29b-41d4-a716-446655440000", false},
		{"hyphen_at_end", "550e8400-e29b-41d4-a716-446655440000-", false},

		// Invalid UUIDs - Invalid characters
		{"with_space", "550e8400 e29b-41d4-a716-446655440000", false},
		{"with_g", "550e8400-e29b-41d4-g716-446655440000", false},
		{"with_z", "z50e8400-e29b-41d4-a716-446655440000", false},
		{"with_special_char", "550e8400-e29b-41d4-a716-44665544000!", false},
		{"with_underscore", "550e8400_e29b_41d4_a716_446655440000", false},
		{"with_plus", "550e8400+e29b+41d4+a716+446655440000", false},

		// Invalid UUIDs - Wrong version
		{"version_0", "550e8400-e29b-01d4-a716-446655440000", false},
		{"version_6", "550e8400-e29b-61d4-a716-446655440000", false},
		{"version_7", "550e8400-e29b-71d4-a716-446655440000", false},
		{"version_9", "550e8400-e29b-91d4-a716-446655440000", false},
		{"version_letter", "550e8400-e29b-a1d4-a716-446655440000", false},

		// Invalid UUIDs - Wrong variant
		{"variant_0", "550e8400-e29b-41d4-0716-446655440000", false},
		{"variant_1", "550e8400-e29b-41d4-1716-446655440000", false},
		{"variant_c", "550e8400-e29b-41d4-c716-446655440000", false},
		{"variant_d", "550e8400-e29b-41d4-d716-446655440000", false},
		{"variant_e", "550e8400-e29b-41d4-e716-446655440000", false},
		{"variant_f", "550e8400-e29b-41d4-f716-446655440000", false},
		{"variant_lowercase_c", "550e8400-e29b-41d4-c716-446655440000", false},
		{"variant_uppercase_c", "550e8400-e29b-41d4-C716-446655440000", false},

		// Invalid UUIDs - Braces format (not supported)
		{"with_braces", "{550e8400-e29b-41d4-a716-446655440000}", false},
		{"urn_format", "urn:uuid:550e8400-e29b-41d4-a716-446655440000", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidUUID(tt.uuid)
			if result != tt.expected {
				t.Errorf("IsValidUUID(%q) = %v, expected %v", tt.uuid, result, tt.expected)
			}
		})
	}
}

func TestIsValidHexChar(t *testing.T) {
	tests := []struct {
		name     string
		char     byte
		expected bool
	}{
		// Valid hex characters
		{"digit_0", '0', true},
		{"digit_5", '5', true},
		{"digit_9", '9', true},
		{"lowercase_a", 'a', true},
		{"lowercase_d", 'd', true},
		{"lowercase_f", 'f', true},
		{"uppercase_A", 'A', true},
		{"uppercase_D", 'D', true},
		{"uppercase_F", 'F', true},

		// Invalid characters
		{"lowercase_g", 'g', false},
		{"lowercase_z", 'z', false},
		{"uppercase_G", 'G', false},
		{"uppercase_Z", 'Z', false},
		{"space", ' ', false},
		{"hyphen", '-', false},
		{"underscore", '_', false},
		{"plus", '+', false},
		{"special_char", '!', false},
		{"bracket", '[', false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidHexChar(tt.char)
			if result != tt.expected {
				t.Errorf("isValidHexChar(%q) = %v, expected %v", tt.char, result, tt.expected)
			}
		})
	}
}

func TestHasValidHyphens(t *testing.T) {
	tests := []struct {
		name     string
		uuid     string
		expected bool
	}{
		// Valid hyphen positions
		{"correct_positions", "550e8400-e29b-41d4-a716-446655440000", true},
		{"correct_with_zeros", "00000000-0000-0000-0000-000000000000", true},
		{"correct_with_letters", "abcdefab-abcd-abcd-abcd-abcdefabcdef", true},

		// Invalid hyphen positions
		{"no_hyphens", "550e8400e29b41d4a716446655440000xxxx", false},
		{"missing_first", "550e8400xe29b-41d4-a716-446655440000", false},
		{"missing_second", "550e8400-e29bx41d4-a716-446655440000", false},
		{"missing_third", "550e8400-e29b-41d4xa716-446655440000", false},
		{"missing_fourth", "550e8400-e29b-41d4-a716x446655440000", false},
		{"wrong_char_1", "550e8400_e29b-41d4-a716-446655440000", false},
		{"wrong_char_2", "550e8400-e29b_41d4-a716-446655440000", false},
		{"wrong_char_3", "550e8400-e29b-41d4_a716-446655440000", false},
		{"wrong_char_4", "550e8400-e29b-41d4-a716_446655440000", false},

		// Edge cases - need exactly 36 chars for valid test
		{"all_hyphens_36", "--------x----x----x----x------------", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := hasValidHyphens(tt.uuid)
			if result != tt.expected {
				t.Errorf("hasValidHyphens(%q) = %v, expected %v", tt.uuid, result, tt.expected)
			}
		})
	}
}

func TestHasValidHexChars(t *testing.T) {
	tests := []struct {
		name     string
		uuid     string
		expected bool
	}{
		// Valid hex characters
		{"all_valid_lowercase", "550e8400-e29b-41d4-a716-446655440000", true},
		{"all_valid_uppercase", "550E8400-E29B-41D4-A716-446655440000", true},
		{"all_valid_mixed", "550e8400-E29B-41d4-A716-446655440000", true},
		{"all_numbers", "12345678-1234-1234-1234-123456789012", true},
		{"all_valid_letters", "abcdefab-abcd-abcd-abcd-abcdefabcdef", true},

		// Invalid characters in different positions
		{"invalid_at_start", "g50e8400-e29b-41d4-a716-446655440000", false},
		{"invalid_in_first_group", "550e840g-e29b-41d4-a716-446655440000", false},
		{"invalid_in_second_group", "550e8400-e2gb-41d4-a716-446655440000", false},
		{"invalid_in_third_group", "550e8400-e29b-41g4-a716-446655440000", false},
		{"invalid_in_fourth_group", "550e8400-e29b-41d4-a7g6-446655440000", false},
		{"invalid_in_last_group", "550e8400-e29b-41d4-a716-44665544000g", false},
		{"space_character", "550e8400-e29b-41d4-a716-44665544000 ", false},
		{"special_character", "550e8400-e29b-41d4-a716-44665544000!", false},
		{"underscore", "550e8400-e29b-41d4-a716-44665544000_", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := hasValidHexChars(tt.uuid)
			if result != tt.expected {
				t.Errorf("hasValidHexChars(%q) = %v, expected %v", tt.uuid, result, tt.expected)
			}
		})
	}
}

func TestIsValidUUIDVersionAndVariant(t *testing.T) {
	tests := []struct {
		name     string
		uuid     string
		expected bool
	}{
		// Valid version 1-5 with valid variants
		{"v1_variant_8", "550e8400-e29b-11d4-8716-446655440000", true},
		{"v1_variant_9", "550e8400-e29b-11d4-9716-446655440000", true},
		{"v1_variant_a", "550e8400-e29b-11d4-a716-446655440000", true},
		{"v1_variant_b", "550e8400-e29b-11d4-b716-446655440000", true},
		{"v2_variant_8", "550e8400-e29b-21d4-8716-446655440000", true},
		{"v3_variant_9", "550e8400-e29b-31d4-9716-446655440000", true},
		{"v4_variant_a", "550e8400-e29b-41d4-a716-446655440000", true},
		{"v5_variant_b", "550e8400-e29b-51d4-b716-446655440000", true},

		// Valid with case variations
		{"v4_variant_uppercase_A", "550e8400-e29b-41d4-A716-446655440000", true},
		{"v4_variant_uppercase_B", "550e8400-e29b-41d4-B716-446655440000", true},
		{"v4_variant_lowercase_a", "550e8400-e29b-41d4-a716-446655440000", true},
		{"v4_variant_lowercase_b", "550e8400-e29b-41d4-b716-446655440000", true},

		// Invalid versions
		{"version_0", "550e8400-e29b-01d4-a716-446655440000", false},
		{"version_6", "550e8400-e29b-61d4-a716-446655440000", false},
		{"version_7", "550e8400-e29b-71d4-a716-446655440000", false},
		{"version_8", "550e8400-e29b-81d4-a716-446655440000", false},
		{"version_9", "550e8400-e29b-91d4-a716-446655440000", false},
		{"version_a", "550e8400-e29b-a1d4-a716-446655440000", false},
		{"version_f", "550e8400-e29b-f1d4-a716-446655440000", false},

		// Invalid variants
		{"variant_0", "550e8400-e29b-41d4-0716-446655440000", false},
		{"variant_1", "550e8400-e29b-41d4-1716-446655440000", false},
		{"variant_2", "550e8400-e29b-41d4-2716-446655440000", false},
		{"variant_3", "550e8400-e29b-41d4-3716-446655440000", false},
		{"variant_4", "550e8400-e29b-41d4-4716-446655440000", false},
		{"variant_5", "550e8400-e29b-41d4-5716-446655440000", false},
		{"variant_6", "550e8400-e29b-41d4-6716-446655440000", false},
		{"variant_7", "550e8400-e29b-41d4-7716-446655440000", false},
		{"variant_c", "550e8400-e29b-41d4-c716-446655440000", false},
		{"variant_d", "550e8400-e29b-41d4-d716-446655440000", false},
		{"variant_e", "550e8400-e29b-41d4-e716-446655440000", false},
		{"variant_f", "550e8400-e29b-41d4-f716-446655440000", false},
		{"variant_uppercase_C", "550e8400-e29b-41d4-C716-446655440000", false},
		{"variant_uppercase_D", "550e8400-e29b-41d4-D716-446655440000", false},

		// Invalid version with valid variant
		{"v0_valid_variant", "550e8400-e29b-01d4-a716-446655440000", false},
		{"v6_valid_variant", "550e8400-e29b-61d4-a716-446655440000", false},

		// Valid version with invalid variant
		{"v4_invalid_variant", "550e8400-e29b-41d4-c716-446655440000", false},
		{"v5_invalid_variant", "550e8400-e29b-51d4-0716-446655440000", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isValidUUIDVersionAndVariant(tt.uuid)
			if result != tt.expected {
				t.Errorf("isValidUUIDVersionAndVariant(%q) = %v, expected %v", tt.uuid, result, tt.expected)
			}
		})
	}
}

// FuzzIsValidUUID performs fuzz testing on UUID validation.
func FuzzIsValidUUID(f *testing.F) {
	addUUIDFuzzSeeds(f)

	f.Fuzz(func(t *testing.T, uuid string) {
		// The function should never panic, regardless of input
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("IsValidUUID panicked on input %q: %v", uuid, r)
			}
		}()

		result := IsValidUUID(uuid)

		// Basic invariants that should always hold
		if result {
			validateUUIDStructure(t, uuid)
		}

		// Test that the function is deterministic
		result2 := IsValidUUID(uuid)
		if result != result2 {
			t.Errorf("IsValidUUID(%q) is not deterministic: got %v then %v", uuid, result, result2)
		}
	})
}

func addUUIDFuzzSeeds(f *testing.F) {
	f.Helper()

	seeds := []string{
		"550e8400-e29b-41d4-a716-446655440000",
		"f47ac10b-58cc-4372-a567-0e02b2c3d479",
		"00000000-0000-0000-0000-000000000000",
		"ffffffff-ffff-ffff-ffff-ffffffffffff",
		"550e8400-e29b-11d4-8716-446655440000",
		"550e8400-e29b-21d4-9716-446655440000",
		"550e8400-e29b-31d4-a716-446655440000",
		"550e8400-e29b-51d4-b716-446655440000",
		"550E8400-E29B-41D4-A716-446655440000",
		"550e8400e29b41d4a716446655440000",
		"550e8400-e29b-41d4-a716-44665544000",
		"550e8400-e29b-41d4-a716-4466554400000",
		"550e8400-e29b-41d4-g716-446655440000",
		"550e8400_e29b_41d4_a716_446655440000",
		"{550e8400-e29b-41d4-a716-446655440000}",
		"550e8400-e29b-61d4-a716-446655440000",
		"550e8400-e29b-41d4-c716-446655440000",
		"550e8400-e29b-41d4-0716-446655440000",
		"",
		"550e8400-e29b-41d4-a716-446655440000-extra",
		"550e8400 e29b 41d4 a716 446655440000",
		"550e8400+e29b+41d4+a716+446655440000",
		"550e8400.e29b.41d4.a716.446655440000",
	}

	for _, seed := range seeds {
		f.Add(seed)
	}
}

func validateUUIDStructure(t *testing.T, uuid string) {
	t.Helper()

	// Valid UUIDs must be exactly 36 characters
	if len(uuid) != 36 {
		t.Errorf("IsValidUUID(%q) returned true but length is %d, expected 36", uuid, len(uuid))

		return
	}

	validateUUIDHyphens(t, uuid)
	validateUUIDVersionAndVariant(t, uuid)
	validateUUIDHexChars(t, uuid)
}

func validateUUIDHyphens(t *testing.T, uuid string) {
	t.Helper()

	// Must have hyphens at positions 8, 13, 18, 23
	if len(uuid) >= 24 {
		expectedHyphenPositions := []int{8, 13, 18, 23}
		for _, pos := range expectedHyphenPositions {
			if uuid[pos] != '-' {
				t.Errorf("IsValidUUID(%q) returned true but missing hyphen at position %d", uuid, pos)
			}
		}
	}
}

func validateUUIDVersionAndVariant(t *testing.T, uuid string) {
	t.Helper()

	// Special cases: allow nil UUID and max UUID
	if uuid == "00000000-0000-0000-0000-000000000000" || uuid == "ffffffff-ffff-ffff-ffff-ffffffffffff" {
		return
	}

	validateUUIDVersion(t, uuid)
	validateUUIDVariant(t, uuid)
}

func validateUUIDVersion(t *testing.T, uuid string) {
	t.Helper()

	// Version must be 1-5 (position 14)
	if len(uuid) > 14 {
		version := uuid[14]
		if version < '1' || version > '5' {
			t.Errorf("IsValidUUID(%q) returned true but version is %c, expected 1-5", uuid, version)
		}
	}
}

func validateUUIDVariant(t *testing.T, uuid string) {
	t.Helper()

	// Variant must be 8, 9, A, B (position 19)
	if len(uuid) > 19 {
		variant := uuid[19]
		validVariants := []byte{'8', '9', 'A', 'a', 'B', 'b'}
		isValidVariant := false

		for _, v := range validVariants {
			if variant == v {
				isValidVariant = true

				break
			}
		}

		if !isValidVariant {
			t.Errorf("IsValidUUID(%q) returned true but variant is %c, expected 8,9,A,B", uuid, variant)
		}
	}
}

func validateUUIDHexChars(t *testing.T, uuid string) {
	t.Helper()

	// All non-hyphen characters must be valid hex
	if len(uuid) != 36 {
		return
	}

	for i, c := range uuid {
		if isHyphenPosition(i) {
			continue // skip hyphens
		}

		if !isValidHexChar(byte(c)) {
			t.Errorf("IsValidUUID(%q) returned true but contains invalid hex char %c at position %d", uuid, c, i)
		}
	}
}

func isHyphenPosition(i int) bool {
	return i == 8 || i == 13 || i == 18 || i == 23
}
