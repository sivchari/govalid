// Package validationhelper provides validation helper functions for govalid.
package validationhelper

// IsValidUUID validates UUID format manually for maximum performance.
// This function implements UUID validation according to RFC 4122 specification.
//
// UUID format: XXXXXXXX-XXXX-MXXX-NXXX-XXXXXXXXXXXX
// Where:
//
//	X = any hexadecimal digit (0-9, a-f, A-F)
//	M = version field (1-5)
//	N = variant field (8, 9, A, B)
//
// Structure (36 characters total):
//
//	8 hex digits  - time_low
//	4 hex digits  - time_mid
//	4 hex digits  - time_hi_and_version (version in 3rd position)
//	4 hex digits  - clock_seq_hi_and_reserved (variant in 1st position)
//	12 hex digits - node
//
// Example breakdown:
//
//	550e8400-e29b-41d4-a716-446655440000
//	^^^^^^^^|^^^^|^^^^|^^^^|^^^^^^^^^^^^^
//	8       |4   |4   |4   |12 hex digits
//	time_low|    |ver |var |node
//	        |time|    |clk |
//	        |_mid|    |_seq|
//
// Position indices:
//
//	0-7:   time_low
//	8:     hyphen
//	9-12:  time_mid
//	13:    hyphen
//	14-17: time_hi_and_version (version at position 14)
//	18:    hyphen
//	19-22: clock_seq (variant at position 19)
//	23:    hyphen
//	24-35: node
//
// Valid versions (RFC 4122):
//
//	1 = Time-based
//	2 = DCE Security
//	3 = Name-based (MD5)
//	4 = Random
//	5 = Name-based (SHA-1)
//
// Valid variants (high nibble of clock_seq_hi):
//
//	8, 9, A, B = RFC 4122 variant (10xx in binary)
func IsValidUUID(s string) bool {
	// Step 1: Check length - must be exactly 36 characters
	// 32 hex digits + 4 hyphens = 36 total
	if len(s) != 36 {
		return false
	}

	// Step 2: Check hyphen positions
	// Hyphens must be at positions 8, 13, 18, 23 to form 8-4-4-4-12 pattern
	if !hasValidHyphens(s) {
		return false
	}

	// Step 3: Check that all non-hyphen characters are valid hexadecimal
	// Valid hex: 0-9, a-f, A-F
	if !hasValidHexChars(s) {
		return false
	}

	// Step 4: Validate version and variant fields according to RFC 4122
	return isValidUUIDVersionAndVariant(s)
}

// isValidHexChar checks if a character is a valid hexadecimal digit.
// Valid hex characters are:
//
//	0-9 (ASCII 48-57)
//	a-f (ASCII 97-102)
//	A-F (ASCII 65-70)
func isValidHexChar(c byte) bool {
	return (c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')
}

// hasValidHyphens checks if hyphens are in correct positions.
// UUID format requires hyphens at specific positions to separate the fields:
//
// Position:  0        8  9      13 14     18 19     23 24           35
// Format:    XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX
// Fields:    time_low |mid |ver  |var  |node
//
// Example:   550e8400-e29b-41d4-a716-446655440000
//
//	^^^^^^^^ ^^^^ ^^^^ ^^^^ ^^^^^^^^^^^^
//
// Positions: 01234567 8    13   18   23.
func hasValidHyphens(s string) bool {
	return s[8] == '-' && s[13] == '-' && s[18] == '-' && s[23] == '-'
}

// hasValidHexChars checks if all non-hyphen characters are valid hex.
// Iterates through all 36 positions, skipping the 4 hyphen positions.
//
// Character validation by position:
//
//	0-7:   hex digits (time_low)
//	8:     hyphen (skip)
//	9-12:  hex digits (time_mid)
//	13:    hyphen (skip)
//	14-17: hex digits (time_hi_and_version)
//	18:    hyphen (skip)
//	19-22: hex digits (clock_seq)
//	23:    hyphen (skip)
//	24-35: hex digits (node)
func hasValidHexChars(s string) bool {
	for i := range 36 {
		// Skip hyphen positions
		if i == 8 || i == 13 || i == 18 || i == 23 {
			continue
		}

		if !isValidHexChar(s[i]) {
			return false
		}
	}

	return true
}

// isValidUUIDVersionAndVariant checks version and variant fields.
// According to RFC 4122, UUIDs have specific requirements for these fields.
//
// Version field (position 14 - 3rd character of 3rd group):
//
//	The version number is in the most significant 4 bits of the
//	time_hi_and_version field. Valid versions are 1-5.
//
// Variant field (position 19 - 1st character of 4th group):
//
//	The variant field determines the layout of the UUID.
//	RFC 4122 specifies variant 10xx (binary), which in hex is:
//	8 (1000), 9 (1001), A (1010), B (1011)
//
// Special cases:
//
//	Nil UUID (all zeros) is valid: 00000000-0000-0000-0000-000000000000
//	Max UUID (all f's) is valid: ffffffff-ffff-ffff-ffff-ffffffffffff
//
// Example with version 4 and variant A:
//
//	550e8400-e29b-41d4-a716-446655440000
//	               ^   ^
//	               |   variant (position 19)
//	               version (position 14)
func isValidUUIDVersionAndVariant(s string) bool {
	// Special case: nil UUID (all zeros) is valid
	if s == "00000000-0000-0000-0000-000000000000" {
		return true
	}

	// Special case: max UUID (all f's) is valid
	if s == "ffffffff-ffff-ffff-ffff-ffffffffffff" {
		return true
	}

	// Check version (position 14): must be 1-5
	// Version represents the UUID generation algorithm:
	//   1 = Time-based (MAC address + timestamp)
	//   2 = DCE Security (POSIX UID/GID + timestamp)
	//   3 = Name-based using MD5 hashing
	//   4 = Random or pseudo-random
	//   5 = Name-based using SHA-1 hashing
	version := s[14]
	if version < '1' || version > '5' {
		return false
	}

	// Check variant (position 19): must be 8, 9, A, B (case insensitive)
	// The variant bits are the 2 most significant bits of clock_seq_hi_and_reserved.
	// RFC 4122 variant (10xx in binary) maps to these hex values:
	//   8 = 1000 (binary) - variant bits: 10
	//   9 = 1001 (binary) - variant bits: 10
	//   A = 1010 (binary) - variant bits: 10
	//   B = 1011 (binary) - variant bits: 10
	// Other variants (0xxx, 110x, 111x) are reserved for other UUID standards.
	variant := s[19]

	return variant == '8' || variant == '9' ||
		variant == 'A' || variant == 'a' ||
		variant == 'B' || variant == 'b'
}
