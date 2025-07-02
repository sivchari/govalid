package unit

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sivchari/govalid/test"
)

func TestMaxItemsValidation(t *testing.T) {
	validate := validator.New()

	tests := []struct {
		name        string
		data        test.MaxItems
		expectError bool
	}{
		{
			name:        "empty slice",
			data:        test.MaxItems{Items: []string{}},
			expectError: false,
		},
		{
			name:        "slice limit minus one",
			data:        test.MaxItems{Items: []string{"a", "b", "c", "d"}},
			expectError: false,
		},
		{
			name:        "slice exactly at limit",
			data:        test.MaxItems{Items: []string{"a", "b", "c", "d", "e"}},
			expectError: false,
		},
		{
			name:        "slice limit plus one",
			data:        test.MaxItems{Items: []string{"a", "b", "c", "d", "e", "f"}},
			expectError: true,
		},
		{
			name:        "nil slice",
			data:        test.MaxItems{Items: nil},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test govalid
			govalidErr := test.ValidateMaxItems(&tt.data)
			govalidHasError := govalidErr != nil

			// Test go-playground/validator
			playgroundErr := validate.Struct(&tt.data)
			playgroundHasError := playgroundErr != nil

			// Compare results
			if govalidHasError != tt.expectError {
				t.Errorf("govalid: expected error=%v, got error=%v (%v)", tt.expectError, govalidHasError, govalidErr)
			}
			if playgroundHasError != tt.expectError {
				t.Errorf("go-playground: expected error=%v, got error=%v (%v)", tt.expectError, playgroundHasError, playgroundErr)
			}
			if govalidHasError != playgroundHasError {
				t.Errorf("behavior mismatch: govalid=%v, playground=%v", govalidHasError, playgroundHasError)
			}
		})
	}
}

// TestMaxItemsMapValidation tests map validation separately since go-playground/validator doesn't support it
func TestMaxItemsMapValidation(t *testing.T) {
	tests := []struct {
		name        string
		data        test.MaxItems
		expectError bool
	}{
		{
			name:        "map within limit",
			data:        test.MaxItems{Items: []string{"a"}, Metadata: map[string]string{"key1": "val1", "key2": "val2"}},
			expectError: false,
		},
		{
			name:        "map exactly at limit",
			data:        test.MaxItems{Items: []string{"a"}, Metadata: map[string]string{"key1": "val1", "key2": "val2", "key3": "val3"}},
			expectError: false,
		},
		{
			name:        "map exceeds limit",
			data:        test.MaxItems{Items: []string{"a"}, Metadata: map[string]string{"key1": "val1", "key2": "val2", "key3": "val3", "key4": "val4"}},
			expectError: true,
		},
		{
			name:        "nil map",
			data:        test.MaxItems{Items: []string{"a"}, Metadata: nil},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test govalid only (go-playground doesn't support map length validation)
			govalidErr := test.ValidateMaxItems(&tt.data)
			govalidHasError := govalidErr != nil

			if govalidHasError != tt.expectError {
				t.Errorf("govalid: expected error=%v, got error=%v (%v)", tt.expectError, govalidHasError, govalidErr)
			}
		})
	}
}

// TestMaxItemsChanValidation tests channel validation separately since go-playground/validator doesn't support it
func TestMaxItemsChanValidation(t *testing.T) {
	tests := []struct {
		name        string
		setupChan   func() chan int
		expectError bool
	}{
		{
			name: "chan within limit",
			setupChan: func() chan int {
				ch := make(chan int, 1)
				ch <- 1
				return ch
			},
			expectError: false,
		},
		{
			name: "chan exactly at limit",
			setupChan: func() chan int {
				ch := make(chan int, 2)
				ch <- 1
				ch <- 2
				return ch
			},
			expectError: false,
		},
		{
			name: "chan exceeds limit",
			setupChan: func() chan int {
				ch := make(chan int, 5)
				ch <- 1
				ch <- 2
				ch <- 3
				return ch
			},
			expectError: true,
		},
		{
			name: "nil chan",
			setupChan: func() chan int {
				return nil
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test govalid only (go-playground doesn't support channel validation)
			data := test.MaxItems{Items: []string{"a"}, ChanField: tt.setupChan()}
			govalidErr := test.ValidateMaxItems(&data)
			govalidHasError := govalidErr != nil

			if govalidHasError != tt.expectError {
				t.Errorf("govalid: expected error=%v, got error=%v (%v)", tt.expectError, govalidHasError, govalidErr)
			}
		})
	}
}
