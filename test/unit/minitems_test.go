package unit

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sivchari/govalid/test"
)

func TestMinItemsSliceValidation(t *testing.T) {
	validate := validator.New()

	tests := []struct {
		name        string
		data        test.MinItems
		expectError bool
	}{
		{
			name:        "empty slice",
			data:        test.MinItems{Items: []string{}, Metadata: map[string]string{"key": "val"}, ChanField: func() chan int { ch := make(chan int, 1); ch <- 1; return ch }()},
			expectError: true,
		},
		{
			name:        "slice limit minus one",
			data:        test.MinItems{Items: []string{"a"}, Metadata: map[string]string{"key": "val"}, ChanField: func() chan int { ch := make(chan int, 1); ch <- 1; return ch }()},
			expectError: true,
		},
		{
			name:        "slice exactly at limit",
			data:        test.MinItems{Items: []string{"a", "b"}, Metadata: map[string]string{"key": "val"}, ChanField: func() chan int { ch := make(chan int, 1); ch <- 1; return ch }()},
			expectError: false,
		},
		{
			name:        "slice limit plus one",
			data:        test.MinItems{Items: []string{"a", "b", "c"}, Metadata: map[string]string{"key": "val"}, ChanField: func() chan int { ch := make(chan int, 1); ch <- 1; return ch }()},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test govalid
			govalidErr := test.ValidateMinItems(&tt.data)
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

// TestMinItemsMapValidation tests map validation separately since go-playground/validator doesn't support it
func TestMinItemsMapValidation(t *testing.T) {
	tests := []struct {
		name        string
		data        test.MinItems
		expectError bool
	}{
		{
			name:        "map below limit",
			data:        test.MinItems{Items: []string{"a", "b"}, Metadata: map[string]string{}, ChanField: func() chan int { ch := make(chan int, 1); ch <- 1; return ch }()},
			expectError: true,
		},
		{
			name:        "map exactly at limit",
			data:        test.MinItems{Items: []string{"a", "b"}, Metadata: map[string]string{"key1": "val1"}, ChanField: func() chan int { ch := make(chan int, 1); ch <- 1; return ch }()},
			expectError: false,
		},
		{
			name:        "map above limit",
			data:        test.MinItems{Items: []string{"a", "b"}, Metadata: map[string]string{"key1": "val1", "key2": "val2"}, ChanField: func() chan int { ch := make(chan int, 1); ch <- 1; return ch }()},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test govalid only (go-playground doesn't support map length validation)
			govalidErr := test.ValidateMinItems(&tt.data)
			govalidHasError := govalidErr != nil

			if govalidHasError != tt.expectError {
				t.Errorf("govalid: expected error=%v, got error=%v (%v)", tt.expectError, govalidHasError, govalidErr)
			}
		})
	}
}

// TestMinItemsChanValidation tests channel validation separately since go-playground/validator doesn't support it
func TestMinItemsChanValidation(t *testing.T) {
	tests := []struct {
		name        string
		setupChan   func() chan int
		expectError bool
	}{
		{
			name: "chan below limit",
			setupChan: func() chan int {
				ch := make(chan int, 1)
				return ch
			},
			expectError: true,
		},
		{
			name: "chan exactly at limit",
			setupChan: func() chan int {
				ch := make(chan int, 1)
				ch <- 1
				return ch
			},
			expectError: false,
		},
		{
			name: "chan above limit",
			setupChan: func() chan int {
				ch := make(chan int, 5)
				ch <- 1
				ch <- 2
				return ch
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test govalid only (go-playground doesn't support channel validation)
			data := test.MinItems{Items: []string{"a", "b"}, Metadata: map[string]string{"key": "val"}, ChanField: tt.setupChan()}
			govalidErr := test.ValidateMinItems(&data)
			govalidHasError := govalidErr != nil

			if govalidHasError != tt.expectError {
				t.Errorf("govalid: expected error=%v, got error=%v (%v)", tt.expectError, govalidHasError, govalidErr)
			}
		})
	}
}
