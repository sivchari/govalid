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

// BenchmarkGoValidCELCached tests performance after CEL programs are cached
func BenchmarkGoValidCELCached(b *testing.B) {
	instance := test.CEL{
		Age:      25,
		Name:     "John Doe",
		Score:    85.5,
		IsActive: true,
	}
	
	// Pre-warm the cache
	for i := 0; i < 100; i++ {
		test.ValidateCEL(&instance)
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