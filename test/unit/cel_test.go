package unit

import (
	"testing"

	"github.com/sivchari/govalid/test"
)

func TestCELValidation(t *testing.T) {
	tests := []struct {
		name        string
		data        test.CEL
		expectError bool
	}{
		{
			name: "valid score",
			data: test.CEL{
				Age:      25,
				Name:     "John",
				Score:    85.5,
				IsActive: true,
			},
			expectError: false,
		},
		{
			name: "zero score",
			data: test.CEL{
				Age:      25,
				Name:     "John",
				Score:    0.0,
				IsActive: true,
			},
			expectError: true,
		},
		{
			name: "negative score",
			data: test.CEL{
				Age:      25,
				Name:     "John",
				Score:    -10.5,
				IsActive: true,
			},
			expectError: true,
		},
		{
			name: "positive score boundary",
			data: test.CEL{
				Age:      25,
				Name:     "John",
				Score:    0.1,
				IsActive: true,
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test govalid CEL validation
			govalidErr := test.ValidateCEL(&tt.data)
			govalidHasError := govalidErr != nil

			// Compare results
			if govalidHasError != tt.expectError {
				t.Errorf("govalid CEL: expected error=%v, got error=%v (%v)", tt.expectError, govalidHasError, govalidErr)
			}

			// Log for debugging
			if govalidHasError {
				t.Logf("CEL validation error: %v", govalidErr)
			} else {
				t.Logf("CEL validation passed for score: %f", tt.data.Score)
			}
		})
	}
}

func TestCELCrossFieldValidation(t *testing.T) {
	tests := []struct {
		name        string
		data        test.CELCrossField
		expectError bool
	}{
		{
			name: "valid price and quantity",
			data: test.CELCrossField{
				Price:    99.99,
				MaxPrice: 150.0,
				Quantity: 2.0,
				Budget:   500.0,
			},
			expectError: false,
		},
		{
			name: "price exceeds max",
			data: test.CELCrossField{
				Price:    200.0,
				MaxPrice: 150.0,
				Quantity: 2.0,
				Budget:   500.0,
			},
			expectError: true,
		},
		{
			name: "quantity * price exceeds budget",
			data: test.CELCrossField{
				Price:    99.99,
				MaxPrice: 150.0,
				Quantity: 6.0,
				Budget:   500.0,
			},
			expectError: true,
		},
		{
			name: "boundary case - price equals max",
			data: test.CELCrossField{
				Price:    150.0,
				MaxPrice: 150.0,
				Quantity: 2.0,
				Budget:   500.0,
			},
			expectError: true,
		},
		{
			name: "boundary case - total equals budget",
			data: test.CELCrossField{
				Price:    100.0,
				MaxPrice: 150.0,
				Quantity: 5.0,
				Budget:   500.0,
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test govalid CEL cross-field validation
			govalidErr := test.ValidateCELCrossField(&tt.data)
			govalidHasError := govalidErr != nil

			// Compare results
			if govalidHasError != tt.expectError {
				t.Errorf("govalid CEL cross-field: expected error=%v, got error=%v (%v)", tt.expectError, govalidHasError, govalidErr)
			}

			// Log for debugging
			if govalidHasError {
				t.Logf("CEL cross-field validation error: %v", govalidErr)
			} else {
				t.Logf("CEL cross-field validation passed for price: %f, quantity: %f", tt.data.Price, tt.data.Quantity)
			}
		})
	}
}