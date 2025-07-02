package benchmark

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sivchari/govalid/test"
)

func BenchmarkGoValidMaxItems(b *testing.B) {
	instance := test.MaxItems{
		Items: []string{"item1", "item2", "item3"},
	}
	b.ResetTimer()
	for b.Loop() {
		err := test.ValidateMaxItems(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
	b.StopTimer()
}

func BenchmarkGoPlaygroundMaxItems(b *testing.B) {
	validate := validator.New()
	instance := test.MaxItems{
		Items: []string{"item1", "item2", "item3"},
	}
	b.ResetTimer()
	for b.Loop() {
		err := validate.Struct(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
	b.StopTimer()
}
