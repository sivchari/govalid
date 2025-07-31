package validationhelper

import "testing"

func TestIsNumeric(t *testing.T) {
	t.Run("valid_numeric_strings", func(t *testing.T) {
		validInputs := []string{
			"0",
			"123",
			"999999",
			"0000000001",
			"12345678901234567890",
		}

		for _, input := range validInputs {
			if !IsNumeric(input) {
				t.Errorf("IsNumeric(%q) = false, expected true", input)
			}
		}
	})

	t.Run("invalid_numeric_strings", func(t *testing.T) {
		invalidInputs := []string{
			"",
			" ",
			"abc",
			"123abc",
			"12.34",
			"123 ",
			" 123",
			"123-456",
			"12_34",
			"1,234",
			"1+2",
		}

		for _, input := range invalidInputs {
			if IsNumeric(input) {
				t.Errorf("IsNumeric(%q) = true, expected false", input)
			}
		}
	})
}
