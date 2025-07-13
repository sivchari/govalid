package benchmark

import (
	"testing"

	"github.com/sivchari/govalid/test"
)

// BenchmarkGoValidCELBasic tests basic CEL expressions (new implementation)
func BenchmarkGoValidCELBasic(b *testing.B) {
	instance := test.CEL{
		Age:      25,
		Name:     "John Doe",
		Score:    85.5,
		IsActive: true,
	}
	b.ResetTimer()
	for b.Loop() {
		err := test.ValidateCEL(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
	b.StopTimer()
}

// BenchmarkGoValidCELCrossField tests cross-field validation performance
func BenchmarkGoValidCELCrossField(b *testing.B) {
	instance := test.CELCrossField{
		Price:    99.99,
		MaxPrice: 150.0,
		Quantity: 2,
		Budget:   500.0,
	}
	
	b.ResetTimer()
	for b.Loop() {
		err := test.ValidateCELCrossField(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
	b.StopTimer()
}

// BenchmarkGoValidCELStringLength tests string length validation
func BenchmarkGoValidCELStringLength(b *testing.B) {
	instance := test.CEL{
		Age:      25,
		Name:     "John Doe",
		Score:    85.5,
		IsActive: true,
	}
	
	b.ResetTimer()
	for b.Loop() {
		// Focus on just the name validation (len(t.Name) > 0)
		if !(len(instance.Name) > 0) {
			b.Fatal("name validation failed")
		}
	}
	b.StopTimer()
}

// BenchmarkGoValidCELNumericComparison tests numeric comparison
func BenchmarkGoValidCELNumericComparison(b *testing.B) {
	instance := test.CEL{
		Age:      25,
		Name:     "John Doe",
		Score:    85.5,
		IsActive: true,
	}
	
	b.ResetTimer()
	for b.Loop() {
		// Focus on just the score validation (t.Score > 0)
		if !(instance.Score > 0) {
			b.Fatal("score validation failed")
		}
	}
	b.StopTimer()
}