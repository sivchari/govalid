//go:build test

package unit

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/sivchari/govalid/test"
	"github.com/sivchari/govalid/validation/middleware"
	"github.com/sivchari/govalid/validation/middleware/testfixture"
)

// TODO: Consider generating this list automatically from the test package
// contextValidator wraps any validator for context testing
// to avoid manual updates when new validators are added.
type contextValidator struct {
	name     string
	validate func(context.Context) error
}

// getContextValidators returns all validators that support context cancellation
// This list should be kept in sync with the validators defined in test/marker.go
func getContextValidators() []contextValidator {
	return []contextValidator{
		{
			name: "Required",
			validate: func(ctx context.Context) error {
				return test.ValidateRequiredContext(ctx, &test.Required{
					Name: "John", Age: 25, Items: []string{"item"},
				})
			},
		},
		{
			name: "Alpha",
			validate: func(ctx context.Context) error {
				return test.ValidateAlphaContext(ctx, &test.Alpha{
					FirstName: "John", LastName: "Doe", CountryCode: "US",
				})
			},
		},
		{
			name: "GT",
			validate: func(ctx context.Context) error {
				return test.ValidateGTContext(ctx, &test.GT{Age: 101})
			},
		},
		{
			name: "GTE",
			validate: func(ctx context.Context) error {
				return test.ValidateGTEContext(ctx, &test.GTE{Age: 18})
			},
		},
		{
			name: "LT",
			validate: func(ctx context.Context) error {
				return test.ValidateLTContext(ctx, &test.LT{Age: 5})
			},
		},
		{
			name: "LTE",
			validate: func(ctx context.Context) error {
				return test.ValidateLTEContext(ctx, &test.LTE{Age: 50})
			},
		},
		{
			name: "MaxLength",
			validate: func(ctx context.Context) error {
				return test.ValidateMaxLengthContext(ctx, &test.MaxLength{Name: "short"})
			},
		},
		{
			name: "MinLength",
			validate: func(ctx context.Context) error {
				return test.ValidateMinLengthContext(ctx, &test.MinLength{Name: "longname"})
			},
		},
		{
			name: "Length",
			validate: func(ctx context.Context) error {
				return test.ValidateLengthContext(ctx, &test.Length{Name: "exactly"})
			},
		},
		{
			name: "MaxItems",
			validate: func(ctx context.Context) error {
				return test.ValidateMaxItemsContext(ctx, &test.MaxItems{
					Items: []string{"a", "b"}, Metadata: map[string]string{"k": "v"},
				})
			},
		},
		{
			name: "MinItems",
			validate: func(ctx context.Context) error {
				ch := make(chan int, 1)
				ch <- 1 // Add an item to satisfy minitems=1
				return test.ValidateMinItemsContext(ctx, &test.MinItems{
					Items: []string{"a", "b", "c"}, Metadata: map[string]string{"k": "v"}, ChanField: ch,
				})
			},
		},
		{
			name: "Email",
			validate: func(ctx context.Context) error {
				return test.ValidateEmailContext(ctx, &test.Email{Email: "test@example.com"})
			},
		},
		{
			name: "UUID",
			validate: func(ctx context.Context) error {
				return test.ValidateUUIDContext(ctx, &test.UUID{UUID: "550e8400-e29b-41d4-a716-446655440000"})
			},
		},
		{
			name: "URL",
			validate: func(ctx context.Context) error {
				return test.ValidateURLContext(ctx, &test.URL{URL: "https://example.com"})
			},
		},
		{
			name: "Numeric",
			validate: func(ctx context.Context) error {
				return test.ValidateNumericContext(ctx, &test.Numeric{Number: "12345"})
			},
		},
		{
			name: "Enum",
			validate: func(ctx context.Context) error {
				return test.ValidateEnumContext(ctx, &test.Enum{
					Role: "admin", Level: 2, UserRole: "manager", Priority: 20,
				})
			},
		},
		{
			name: "CEL",
			validate: func(ctx context.Context) error {
				return test.ValidateCELContext(ctx, &test.CEL{
					Age: 25, Name: "John", Score: 85.5, IsActive: true,
				})
			},
		},
		{
			name: "CELCrossField",
			validate: func(ctx context.Context) error {
				return test.ValidateCELCrossFieldContext(ctx, &test.CELCrossField{
					Price: 50.0, MaxPrice: 100.0, Quantity: 2.0, Budget: 200.0,
				})
			},
		},
		{
			name: "IPV4",
			validate: func(ctx context.Context) error {
				return test.ValidateIPV4Context(ctx, &test.IPV4{IP: "192.168.1.1"})
			},
		},
		{
			name: "IPV6",
			validate: func(ctx context.Context) error {
				return test.ValidateIPV6Context(ctx, &test.IPV6{IP: "::1"})
			},
		},
		{
			name: "MultipleErrors",
			validate: func(ctx context.Context) error {
				return test.ValidateMultipleErrorsContext(ctx, &test.MultipleErrors{
					URL: "http://example.com", TooLong: "x",
				})
			},
		},
	}
}

func TestContextCancellation(t *testing.T) {
	t.Parallel()

	validators := getContextValidators()

	for _, v := range validators {
		t.Run(v.name+"_cancellation", func(t *testing.T) {
			t.Parallel()

			ctx, cancel := context.WithCancel(t.Context())
			cancel()

			err := v.validate(ctx)
			if err != context.Canceled {
				t.Errorf("expected context.Canceled, got %v", err)
			}
		})
	}
}

func TestContextTimeout(t *testing.T) {
	t.Parallel()

	validators := getContextValidators()

	for _, v := range validators {
		t.Run(v.name+"_timeout", func(t *testing.T) {
			t.Parallel()

			// TODO: could be improved by using go1.25 synctest package
			ctx, cancel := context.WithTimeout(t.Context(), 1*time.Nanosecond)
			defer cancel()
			time.Sleep(10 * time.Millisecond)

			err := v.validate(ctx)
			if err != context.DeadlineExceeded {
				t.Errorf("expected context.DeadlineExceeded, got %v", err)
			}
		})
	}
}

func TestContextSuccess(t *testing.T) {
	t.Parallel()

	validators := getContextValidators()

	for _, v := range validators {
		t.Run(v.name+"_success", func(t *testing.T) {
			t.Parallel()

			// TODO: could be improved by using go1.25 synctest package
			ctx, cancel := context.WithTimeout(t.Context(), 5*time.Second)
			defer cancel()

			err := v.validate(ctx)
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
		})
	}
}

func TestValidateRequestContext(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		requestBody interface{}
		contextFn   func() (context.Context, context.CancelFunc)
		want        int
	}{
		"valid request with context": {
			requestBody: testfixture.PersonRequest{Name: "John", Email: "john@example.com"},
			contextFn: func() (context.Context, context.CancelFunc) {
				return context.WithTimeout(t.Context(), 5*time.Second)
			},
			want: http.StatusOK,
		},
		"invalid email with context": {
			requestBody: testfixture.PersonRequest{Name: "John", Email: "invalid-email"},
			contextFn: func() (context.Context, context.CancelFunc) {
				return context.WithTimeout(t.Context(), 5*time.Second)
			},
			want: http.StatusBadRequest,
		},
		"cancelled context": {
			requestBody: testfixture.PersonRequest{Name: "John", Email: "john@example.com"},
			contextFn: func() (context.Context, context.CancelFunc) {
				ctx, cancel := context.WithCancel(t.Context())
				cancel()
				return ctx, cancel
			},
			want: http.StatusBadRequest,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			testHandler := func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(http.StatusOK)
			}

			ctx, cancel := tt.contextFn()
			defer cancel()

			jsonBody, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest("POST", "/test", bytes.NewBuffer(jsonBody))
			req = req.WithContext(ctx)
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()
			sut := middleware.ValidateRequestContext[*testfixture.PersonRequest](testHandler)
			sut(rr, req)

			if rr.Code != tt.want {
				t.Errorf("Expected status %d, got %d", tt.want, rr.Code)
			}
		})
	}
}
