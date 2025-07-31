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

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			err := test.ValidateCEL(&instance)
			if err != nil {
				b.Fatal("unexpected error:", err)
			}
		}
	})
}

// BenchmarkGoValidCELMultipleExpressions tests performance with different expressions
func BenchmarkGoValidCELMultipleExpressions(b *testing.B) {
	instance := test.CEL{
		Age:      25,
		Name:     "John",
		Score:    85.5,
		IsActive: true,
	}

	for b.Loop() {
		// Test the same validation multiple times to benefit from caching
		err := test.ValidateCEL(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
}

// BenchmarkGoValidCELCacheEffectiveness measures cache hit performance
func BenchmarkGoValidCELCacheEffectiveness(b *testing.B) {
	var wg sync.WaitGroup
	goroutines := 10

	instance := test.CEL{
		Age:      25,
		Name:     "John Doe",
		Score:    85.5,
		IsActive: true,
	}

	for range goroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for b.Loop() {
				err := test.ValidateCEL(&instance)
				if err != nil {
					b.Error("unexpected error:", err)
					return
				}
			}
		}()
	}
	wg.Wait()
}
