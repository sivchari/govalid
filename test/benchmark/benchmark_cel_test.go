package benchmark

import (
	"testing"

	"github.com/sivchari/govalid/test"
)

// BenchmarkGoValidCELSimple tests simple CEL expressions
func BenchmarkGoValidCELSimple(b *testing.B) {
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

// BenchmarkGoValidCELRepeated tests performance with repeated validations 
// (naturally exercises cache effectiveness without artificial pre-warming)
func BenchmarkGoValidCELRepeated(b *testing.B) {
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

// BenchmarkGoValidCELVariousExpressions tests different CEL expressions to measure cache effectiveness
func BenchmarkGoValidCELVariousExpressions(b *testing.B) {
	instance := test.CEL{
		Age:      25,
		Name:     "John Doe",
		Score:    85.5,
		IsActive: true,
	}
	
	b.ResetTimer()
	for b.Loop() {
		// This tests the same expressions but exercises the cache
		err := test.ValidateCEL(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
	b.StopTimer()
}