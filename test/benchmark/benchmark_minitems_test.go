package benchmark

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sivchari/govalid/test"
)

func BenchmarkGoValidMinItems(b *testing.B) {
	instance := test.MinItems{
		Items:     []string{"item1", "item2", "item3"},
		Metadata:  map[string]string{"key1": "val1", "key2": "val2"},
		ChanField: make(chan int, 2),
	}
	instance.ChanField <- 1
	instance.ChanField <- 2

	b.ResetTimer()
	for b.Loop() {
		err := test.ValidateMinItems(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
	b.StopTimer()
}

func BenchmarkGoPlaygroundMinItems(b *testing.B) {
	validate := validator.New()
	instance := test.MinItems{
		Items:     []string{"item1", "item2", "item3"},
		Metadata:  map[string]string{"key1": "val1", "key2": "val2"},
		ChanField: make(chan int, 2),
	}
	instance.ChanField <- 1
	instance.ChanField <- 2

	b.ResetTimer()
	for b.Loop() {
		err := validate.Struct(&instance)
		if err != nil {
			b.Fatal("unexpected error:", err)
		}
	}
	b.StopTimer()
}
