package benchmark

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sivchari/govalid/test"
)

// customScoreValidation is a custom validation function for go-playground/validator
// that mimics the CEL expression "value > 0.0"
func customScoreValidation(fl validator.FieldLevel) bool {
	score := fl.Field().Float()
	return score > 0.0
}

func BenchmarkGoValidCEL(b *testing.B) {
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

func BenchmarkGoPlaygroundCEL(b *testing.B) {
	// Set up go-playground/validator with custom validation function
	validate := validator.New()
	err := validate.RegisterValidation("score_positive", customScoreValidation)
	if err != nil {
		b.Fatal("failed to register custom validation:", err)
	}

	// Define a struct that mimics our CEL validation with custom tags
	type PlaygroundCEL struct {
		Age      int     `validate:"min=18" json:"age"`
		Name     string  `validate:"required" json:"name"`
		Score    float64 `validate:"score_positive" json:"score"`
		IsActive bool    `json:"is_active"`
	}

	instance := PlaygroundCEL{
		Age:      25,
		Name:     "John Doe",
		Score:    85.5,
		IsActive: true,
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

// BenchmarkGoValidCELComplex tests more complex CEL expressions
func BenchmarkGoValidCELComplex(b *testing.B) {
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

// BenchmarkGoPlaygroundCELComplex tests equivalent complex validation with go-playground
func BenchmarkGoPlaygroundCELComplex(b *testing.B) {
	validate := validator.New()
	
	// Register multiple custom validations to simulate complex CEL expressions
	err := validate.RegisterValidation("score_positive", customScoreValidation)
	if err != nil {
		b.Fatal("failed to register custom validation:", err)
	}

	type ComplexPlaygroundCEL struct {
		Age      int     `validate:"min=18,max=120" json:"age"`
		Name     string  `validate:"required,min=1" json:"name"`
		Score    float64 `validate:"score_positive" json:"score"`
		IsActive bool    `validate:"required" json:"is_active"`
	}

	instance := ComplexPlaygroundCEL{
		Age:      25,
		Name:     "John Doe",
		Score:    85.5,
		IsActive: true,
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