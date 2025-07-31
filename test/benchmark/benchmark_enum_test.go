package benchmark

import (
	"testing"

	"github.com/sivchari/govalid/test"
)

// BenchmarkGoValidEnum tests enum validation performance
// Note: go-playground/validator doesn't have a direct enum equivalent,
// so this benchmark is standalone for govalid
func BenchmarkGoValidEnum(b *testing.B) {
	instance := test.Enum{
		Role:     "admin",
		Level:    1,
		UserRole: "manager",
		Priority: 10,
	}
	for b.Loop() {
		err := test.ValidateEnum(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
}
