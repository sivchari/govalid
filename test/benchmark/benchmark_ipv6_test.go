package benchmark

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sivchari/govalid/test"
)

func BenchmarkGoValidIPV6(b *testing.B) {
	instance := test.IPV6{
		IP: "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
	}
	for b.Loop() {
		err := test.ValidateIPV6(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
}

func BenchmarkGoPlaygroundIPV6(b *testing.B) {
	validate := validator.New()
	instance := test.IPV6{
		IP: "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
	}
	for b.Loop() {
		err := validate.Struct(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
}
