package benchmark

import (
	"sync"
	"testing"

	"github.com/sivchari/govalid/test"
)

// BenchmarkGoValidCELConcurrent tests CEL performance under concurrent load
func BenchmarkGoValidCELConcurrent(b *testing.B) {
	instance := test.CEL{
		Age:      25,
		Name:     "John Doe",
		Score:    85.5,
		IsActive: true,
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			err := test.ValidateCEL(&instance)
			if err != nil {
				b.Fatal("unexpected error:", err)
			}
		}
	})
	b.StopTimer()
}

// BenchmarkGoValidCELMultipleExpressions tests performance with different expressions
func BenchmarkGoValidCELMultipleExpressions(b *testing.B) {
	instance := test.CEL{
		Age:      25,
		Name:     "John",
		Score:    85.5,
		IsActive: true,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Test the same validation multiple times to benefit from caching
		err := test.ValidateCEL(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
	b.StopTimer()
}

// BenchmarkGoValidCELCacheEffectiveness measures cache hit performance
func BenchmarkGoValidCELCacheEffectiveness(b *testing.B) {
	var wg sync.WaitGroup
	goroutines := 10
	iterations := b.N / goroutines

	instance := test.CEL{
		Age:      25,
		Name:     "John Doe", 
		Score:    85.5,
		IsActive: true,
	}

	b.ResetTimer()
	for g := 0; g < goroutines; g++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < iterations; i++ {
				err := test.ValidateCEL(&instance)
				if err != nil {
					b.Error("unexpected error:", err)
					return
				}
			}
		}()
	}
	wg.Wait()
	b.StopTimer()
}