//go:build test

package middleware_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sivchari/govalid/validation/middleware"
	"github.com/sivchari/govalid/validation/middleware/testfixture"
)

func TestValidateRequest(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		requestBody interface{}
		want        int
	}{
		"valid request": {
			requestBody: testfixture.PersonRequest{Name: "John", Email: "john@example.com"},
			want:        http.StatusOK,
		},
		"invalid email": {
			requestBody: testfixture.PersonRequest{Name: "John", Email: "invalid-email"},
			want:        http.StatusBadRequest,
		},
		"invalid json": {
			requestBody: "invalid json",
			want:        http.StatusBadRequest,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			testHandler := func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(http.StatusOK)
			}

			var req *http.Request
			if str, ok := tt.requestBody.(string); ok {
				req = httptest.NewRequest("POST", "/test", bytes.NewBufferString(str))
			} else {
				jsonBody, _ := json.Marshal(tt.requestBody)
				req = httptest.NewRequest("POST", "/test", bytes.NewBuffer(jsonBody))
			}

			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()
			sut := middleware.ValidateRequest[*testfixture.PersonRequest](testHandler)
			sut(rr, req)

			if rr.Code != tt.want {
				t.Errorf("Expected status %d, got %d", tt.want, rr.Code)
			}
		})
	}
}
