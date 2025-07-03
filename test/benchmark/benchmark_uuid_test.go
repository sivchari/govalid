package benchmark

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sivchari/govalid/test"
)

func BenchmarkGoValidUUID(b *testing.B) {
	instance := test.UUID{UUID: "6ba7b811-9dad-41d1-80b4-00c04fd430c8"}
	b.ResetTimer()
	for b.Loop() {
		err := test.ValidateUUID(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
	b.StopTimer()
}

func BenchmarkGoPlaygroundUUID(b *testing.B) {
	validate := validator.New()
	instance := test.UUID{UUID: "6ba7b811-9dad-41d1-80b4-00c04fd430c8"}
	b.ResetTimer()
	for b.Loop() {
		err := validate.Struct(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
	b.StopTimer()
}